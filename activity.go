package incidentrenderactivity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-incidentrenderactivity")

// MyActivity is a stub for your Activity implementation
type XMLParserActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XMLParserActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *XMLParserActivity) Metadata() *activity.Metadata {
	return a.metadata
}

//XSD
type Incidentsinfo struct {
	XMLName      xml.Name   `xml:"IncidentsInfo" json:"-"`
	IncidentList []Incident `xml:"Incident" json:"Incident"`
}

type Incident struct {
	XMLName xml.Name `xml:"Incident" json:"-"`
	XCoor   string   `xml:"XCoor" json:"XCoor"`
	YCoor   string   `xml:"YCoor" json:"YCoor"`
	Type    string   `xml:"Type" json:"Type"`
	Message string   `xml:"Message" json:"Message"`
}

// end of XSD

// Eval implements activity.Activity.Eval
func (a *XMLParserActivity) Eval(ctx activity.Context) (done bool, err error) {

	JsonString := ctx.GetInput("data").(string)

	activityLog.Debugf("Json String is : [%s]", JsonString)
	//fmt.Println("Json String is : ", JsonString)

	if len(JsonString) == 0 {
		activityLog.Debugf("value in  the field is empty ")
		fmt.Println("value in  the field is empty ")
		return
	}
	//fmt.Println(" JSON String " + JsonString)

	jdata := Incidentsinfo{}
	err = json.Unmarshal([]byte(JsonString), &jdata)
	xmlData, _ := xml.Marshal(jdata)

	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}

	//fmt.Println(" XML String ")
	//fmt.Println(string(xmlData))

	// Set the output as part of the context
	activityLog.Debugf("Activity has rendered Incident json Successfully")
	//fmt.Println("Activity has rendered Incident json Successfully")

	ctx.SetOutput("output", string(xmlData))

	return true, nil
}
