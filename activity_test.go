package incidentrenderactivity

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("data", `{"Incident":[{"XCoor":"25133.26","YCoor":"43509.02","Type":"1","Message":"(9/3)14:29 Roadworks on SLE (towards BKE) before Mandai Rd Exit. Avoid lane 1."}
	,{"XCoor":"23835.43","YCoor":"45044.879","Type":"0","Message":"(9/3)14:28 Accident on SLE (towards CTE) at Woodlands Ave 12 Exit."}
	,{"XCoor":"35192.801","YCoor":"34073.34","Type":"0","Message":"(9/3)14:26 Accident on PIE (towards Changi Airport) after Paya Lebar. Avoid lane 1."}]}
	`)
	//tc.SetInput("file", "D:/Flogo/xml.jsp")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	output := tc.GetOutput("output")
	assert.Equal(t, output, output)

}
