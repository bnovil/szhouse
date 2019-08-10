package dao

import (
	"encoding/json"
	"crawl/szhouse/projectList/data"
	"testing"
)

func TestInsert(t *testing.T){
	s := `{"Id":41053,"Seq":"1","PreNumber":"恒大时尚慧谷大厦","NumberLink":"/certdetail.aspx?id=41053","Name":"恒大时尚慧谷大厦","NameLink":"projectdetail.aspx?id=41053","Company":"建滔数码发展（深圳）有限公司","District":"宝安","ApproveTime":"2019-07-30"}`
	var b data.ProjectBrief

	json.Unmarshal([]byte(s),&b)

	f := Insert(b)
	if f{
		t.Log("insert succeed")
	}else{
		t.Log("insert failed")
	}
}