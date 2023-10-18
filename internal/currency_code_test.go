package internal

import (
	"encoding/xml"
	"os"
	"testing"
)

func Test_parse_currency_code(t *testing.T) {
	file := `/Users/yan/work/github.com/beanscc/rango/internal/currency_code.xml`
	var iso4217 struct {
		XMLName xml.Name `xml:"ISO_4217"`
		CcyTbl  struct {
			CcyNtry []struct {
				CtryNm     string `xml:"CtryNm"`     // 国家名
				CcyNm      string `xml:"CcyNm"`      // 货币名
				Ccy        string `xml:"Ccy"`        // 货币代码
				CcyNbr     string `xml:"CcyNbr"`     // 货币数字代码
				CcyMnrUnts string `xml:"CcyMnrUnts"` // 精度单位
			} `xml:"CcyNtry"`
		} `xml:"CcyTbl"`
	}

	content, err := os.ReadFile(file)
	if err != nil {
		t.Errorf("os.ReadFile err:%v", err)
		return
	}

	if err := xml.Unmarshal(content, &iso4217); err != nil {
		t.Errorf("xml.Unmarshal err:%v", err)
	}

	for i, v := range iso4217.CcyTbl.CcyNtry {
		t.Logf("%3d, currency:%+v", i, v)
	}
}
