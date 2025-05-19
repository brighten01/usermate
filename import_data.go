package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

type ResponseData struct {
	tags []*Tags
}

type Tags struct {
}
type CourseContent struct {
	RootID          string `json:"root_id"`
	RootName        string `json:"root_name"`
	LeafnodeID      string `json:"leafnode_id"`
	LeafnodeName    string `json:"leafnode_name"`
	ContentGrade    string `json:"content_grade"`
	ContentSubject  string `json:"content_subject"`
	ContentVersion  string `json:"content_version"`
	ContentSemester string `json:"content_semester"`
	IsNewCourse     int    `json:"is_new_course"`
	TaskType        string `json:"_task_type"`
}

type ApiResponse struct {
	ErrorReason string      `json:"error_reason"`
	ErrorMsg    string      `json:"error_msg"`
	MetaData    interface{} `json:"meta_data"`
	TraceID     string      `json:"trace_id"`
	ServerTime  int64       `json:"server_time"`
	Data        Data        `json:"data"`
}

type Data struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	KnowledgeID          string              `json:"knowledge_id"`
	KnowledgeName        string              `json:"knowledge_name"`
	KnowledgeStandardTag string              `json:"knowledge_standard_tag"`
	KnowledgeLevelStats  KnowledgeLevelStats `json:"knowledge_level_stats"`
	ExamFrequencyTag     string              `json:"exam_frequency_tag"`
	IsMistakes           int                 `json:"is_mistakes"`
}

type KnowledgeLevelStats struct {
	DegreeSCnt int `json:"degree_s_cnt"`
	DegreeACnt int `json:"degree_a_cnt"`
	DegreeBCnt int `json:"degree_b_cnt"`
	DegreeCCnt int `json:"degree_c_cnt"`
	AllCnt     int `json:"all_cnt"`
}

type KnowledgeNode struct {
	// 图谱基本信息
	RootID       string `json:"root_id" gorm:"column:root_id;type:varchar(64);comment:知识图谱根节点ID"`
	LeafNodeID   string `json:"leafnode_id" gorm:"column:leafnode_id;type:varchar(64);comment:叶子节点ID"`
	LeafNodeName string `json:"leafnode_name" gorm:"column:leafnode_name;type:varchar(128);comment:叶子节点名称"`
	// 课程内容信息
	ContentGrade    int    `json:"content_grade" gorm:"column:content_grade;type:tinyint;comment:年级(1-12分别对应小一到高三)"`
	ContentSubject  string `json:"content_subject" gorm:"column:content_subject;type:varchar(32);comment:学科名称"`
	ContentVersion  string `json:"content_version" gorm:"column:content_version;type:varchar(32);comment:教材版本"`
	ContentSemester string `json:"content_semester" gorm:"column:content_semester;type:tinyint;comment:学期(1:上学期,2:下学期)"`
	IsNewCourse     int    `json:"is_new_course" gorm:"column:is_new_course;type:tinyint(1);comment:是否新课标(true:是,false:否)"`
	// 考点标签
	ExamFrequencyTag string `json:"exam_frequency_tag" gorm:"-;comment:高频考点标签列表"`
	IsMistakeProne   int    `json:"is_mistakes" gorm:"column:is_mistakes;type:tinyint(1);comment:是否易错知识点"`
	Standard         string `json:"standard"`
}

func export_data_list() []*CourseContent {
	content, err := os.ReadFile("knowledge.json")
	if err != nil {
		fmt.Println(err)
	}
	var courseData []*CourseContent
	err = json.Unmarshal(content, &courseData)
	return courseData
}

type MyReader struct {
	buf      []byte // cont
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead int    // la
}

func get_knowledge_frequnency(url string, requestData *CurriculumRequest) (apiresponse *ApiResponse, err error) {
	content, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Json Serialize failed ", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		fmt.Println("request error ", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do error ", err)
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("status is ", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body error ", err)
		return nil, err
	}

	var tmpData ApiResponse
	err = json.Unmarshal(body, &tmpData)
	if err != nil {
		fmt.Println("json 解析错误", err)
		return nil, err
	}

	return &tmpData, nil
}

type CurriculumRequest struct {
	VersionID    int      `json:"version_id"`
	GradeID      int      `json:"grade_id"`
	SubjectID    int      `json:"subject_id"`
	SemesterID   int      `json:"semester_id"`
	LocationID   string   `json:"location_id"`
	KnowledgeIDs []string `json:"knowledge_ids"`
	RootID       string   `json:"root_id"`
}

type T struct {
	VersionId    int      `json:"version_id"`
	GradeId      int      `json:"grade_id"`
	SubjectId    int      `json:"subject_id"`
	SemesterId   int      `json:"semester_id"`
	LocationId   string   `json:"location_id"`
	KnowledgeIds []string `json:"knowledge_ids"`
	RootId       string   `json:"root_id"`
}

func WriteLog(format string, a ...interface{}) {
	logContent := fmt.Sprintf(format, a)
	filename := "request_error.log"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err ", err)
	}
	_, err = file.WriteString(logContent)
	if err != nil {
		fmt.Println("open file err ", err)
	}

}

func GetIsmistakes(standard string, b, c, s, a int) int {
	mistake := float64(float64(b+c) / float64(s+a))
	//standardInfo := GetClassStandardMap(standard)
	standardInfo := standard
	if standardInfo == "" {
		return 2
	}
	if standardInfo == "理解" && mistake >= 1.2 && mistake <= 4.5 {
		return 1
	}
	if standardInfo == "掌握" && mistake >= 1.0 && mistake <= 5.0 {
		return 1
	}
	if standardInfo == "运用" && mistake >= 1.2 && mistake <= 4.5 {
		return 1
	}

	return 2
}

// 表头
// 图谱ID，节点名称，节点ID，年级，学科，版本，学期，是否新课标，高频考点标签、易错的标签
//
// 1、都没有高频考点标签、课标要求为运用的标签
// 2、都没有易错的标签
func CheckData(url string, v *CourseContent) []*KnowledgeNode {
	VersionID, _ := strconv.Atoi(v.ContentVersion)
	GradeID, _ := strconv.Atoi(v.ContentGrade)
	SubjectID, _ := strconv.Atoi(v.ContentSubject)
	KnowledgeIds := []string{v.LeafnodeID}

	requestData := &CurriculumRequest{
		VersionID:    VersionID,
		GradeID:      GradeID,
		SubjectID:    SubjectID,
		KnowledgeIDs: KnowledgeIds,
		RootID:       v.RootID,
	}
	response, err := get_knowledge_frequnency(url, requestData)
	if err != nil {
		WriteLog("requestData err %v", err)
		return nil
	}

	//考纲 考频 易错
	var knowledgeData []*KnowledgeNode
	for _, info := range response.Data.Tags {
		b := info.KnowledgeLevelStats.DegreeBCnt
		c := info.KnowledgeLevelStats.DegreeCCnt
		s := info.KnowledgeLevelStats.DegreeSCnt
		a := info.KnowledgeLevelStats.DegreeACnt

		// 将ContentGrade转换为int
		contentGrade, _ := strconv.Atoi(v.ContentGrade)

		// 处理考频标签
		examFrequencyTag := info.ExamFrequencyTag
		if examFrequencyTag == "" {
			examFrequencyTag = "非高频" // 设置默认值
		}

		kowledgeInfo := &KnowledgeNode{
			RootID:           v.RootID,
			LeafNodeID:       v.LeafnodeID,
			LeafNodeName:     v.LeafnodeName,
			ContentGrade:     contentGrade,
			ContentSubject:   v.ContentSubject,
			ContentVersion:   v.ContentVersion,
			ContentSemester:  v.ContentSemester,
			IsNewCourse:      v.IsNewCourse,
			Standard:         info.KnowledgeStandardTag,
			ExamFrequencyTag: examFrequencyTag,
			IsMistakeProne:   GetIsmistakes(info.KnowledgeStandardTag, b, c, s, a),
		}

		knowledgeData = append(knowledgeData, kowledgeInfo)
	}
	return knowledgeData
}

// 图谱ID，节点名称，节点ID，年级，学科，版本，学期，是否新课标，高频考点标签、易错的标签
func advancedHeaderWrite(f *excelize.File, sheet string) {
	_ = f.SetCellValue(sheet, "A1", "图谱ID")
	_ = f.SetCellValue(sheet, "B1", "节点名称")
	_ = f.SetCellValue(sheet, "C1", "节点ID")
	_ = f.SetCellValue(sheet, "D1", "年级")
	_ = f.SetCellValue(sheet, "E1", "学科")
	_ = f.SetCellValue(sheet, "F1", "版本")
	_ = f.SetCellValue(sheet, "G1", "学期")
	_ = f.SetCellValue(sheet, "H1", "考纲")
	_ = f.SetCellValue(sheet, "I1", "高频考点")
	_ = f.SetCellValue(sheet, "J1", "易错")

}

func main() {
	courseData := export_data_list()
	f := excelize.NewFile()

	// 创建新的工作表
	sheet := "考频考点"
	_, err := f.NewSheet(sheet)
	if err != nil {
		fmt.Println("创建工作表失败: ", err)
		return
	}

	// 写入表头
	advancedHeaderWrite(f, sheet)

	var wg sync.WaitGroup
	var mu sync.Mutex // 添加互斥锁
	rowCounter := 2   // 从第2行开始写入数据（第1行是表头）

	chunkSize := 1500
	for i := 0; i < len(courseData); i += chunkSize {
		wg.Add(1)
		end := i + chunkSize
		if end > len(courseData) {
			end = len(courseData)
		}

		go func(start, end int) {
			defer wg.Done()

			// 获取当前批次的数据
			batchData := courseData[start:end]
			var url = "https://genie-internal.vdyoo.net/data-engine-service-gray/v1/question/batch_get_knowledge_class_standard"

			// 处理每个课程内容
			for _, v := range batchData {
				responseData := CheckData(url, v)
				// 使用互斥锁保护写入操作
				mu.Lock()
				currentRow := rowCounter
				rowCounter += len(responseData)
				mu.Unlock()

				// 写入数据
				for i, value := range responseData {
					// fmt.Printf("response data is  %#v", value)
					// os.Exit(1)
					var data []string
					data = append(data, value.RootID)
					data = append(data, value.LeafNodeName)
					data = append(data, value.LeafNodeID)
					data = append(data, strconv.Itoa(value.ContentGrade))
					data = append(data, value.ContentSubject)
					data = append(data, value.ContentVersion)
					data = append(data, value.ContentSemester)
					data = append(data, value.Standard)
					data = append(data, value.ExamFrequencyTag)
					data = append(data, strconv.Itoa(value.IsMistakeProne))
					mu.Lock()
					err := f.SetSheetRow(sheet, "A"+strconv.Itoa(currentRow+i), &data)
					mu.Unlock()
					if err != nil {
						WriteLog("excel写入错误: %v", err)
					}
				}
			}
		}(i, end)
	}

	wg.Wait()
	fileName := fmt.Sprintf("exam_frequency_mistake_%s.xlsx", time.Now().Format("20060102150405"))
	err = f.SaveAs(fileName)
	if err != nil {
		fmt.Println("保存Excel文件错误: ", err)
		WriteLog("保存Excel文件错误: %v", err)
	}
}

