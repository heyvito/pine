package pine

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/logrusorgru/aurora"
)

func testTimeProvider() time.Time {
	time, _ := time.Parse("20060102150405", "20081027102708")
	return time
}

const timeValue string = "10:27:08"

var inputChannel = make(chan string)
var lastOutput string

func testOutputProvider(msg string) {
	lastOutput = msg
}

func generateExpectedOutput(module, msg, emoji string, extra *string) string {
	prefix := fmt.Sprintf("%s %s  %s ", aurora.Gray(timeValue), emoji, aurora.Magenta(module))
	ex := ""
	if extra != nil {
		ex = fmt.Sprintf("%s ", aurora.Cyan(*extra))
	}
	return fmt.Sprintf("%s%s%s\n", prefix, ex, msg)
}

func TestMain(m *testing.M) {
	pine.outputProvider = testOutputProvider
	pine.timeProvider = testTimeProvider
	pine.formatProvider = ttyFormatProvider
	os.Exit(m.Run())
}

func TestBasicFinish(t *testing.T) {
	writer := pine.NewWriter("MRAjW")
	writer.Finish("XVlBz %s", "gbaiC")
	exp := generateExpectedOutput("MRAjW", "XVlBz gbaiC", "🏁", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraFinish(t *testing.T) {
	writer := pine.NewWriter("hxKQF")
	writer.FinishExtra("DaFpL", "whTHc %s", "tcuAx")
	extraVal := "DaFpL"
	exp := generateExpectedOutput("hxKQF", "whTHc tcuAx", "🏁", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterFinish(t *testing.T) {
	writer := pine.NewWriter("RsWxP")
	ext := writer.WithExtra("LDnJO")
	ext.Finish("SjFbc %s", "XoEFf")
	extraVal := "LDnJO"
	exp := generateExpectedOutput("RsWxP", "SjFbc XoEFf", "🏁", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicTerminate(t *testing.T) {
	writer := pine.NewWriter("aPEZQ")
	writer.Terminate("bCsNV %s", "lgTeM")
	exp := generateExpectedOutput("aPEZQ", "bCsNV lgTeM", "⛔️", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraTerminate(t *testing.T) {
	writer := pine.NewWriter("JjPjz")
	writer.TerminateExtra("pfRFE", "leQYh %s", "YzRyW")
	extraVal := "pfRFE"
	exp := generateExpectedOutput("JjPjz", "leQYh YzRyW", "⛔️", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterTerminate(t *testing.T) {
	writer := pine.NewWriter("bZRjx")
	ext := writer.WithExtra("Awnwe")
	ext.Terminate("gmota %s", "FetHs")
	extraVal := "Awnwe"
	exp := generateExpectedOutput("bZRjx", "gmota FetHs", "⛔️", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSpawn(t *testing.T) {
	writer := pine.NewWriter("EkXBA")
	writer.Spawn("krBEm %s", "fdzdc")
	exp := generateExpectedOutput("EkXBA", "krBEm fdzdc", "✨", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSpawn(t *testing.T) {
	writer := pine.NewWriter("TCoaN")
	writer.SpawnExtra("atyyi", "kjQZL %s", "CtTMt")
	extraVal := "atyyi"
	exp := generateExpectedOutput("TCoaN", "kjQZL CtTMt", "✨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSpawn(t *testing.T) {
	writer := pine.NewWriter("Jrscc")
	ext := writer.WithExtra("tNswY")
	ext.Spawn("NKARe %s", "KJyiX")
	extraVal := "tNswY"
	exp := generateExpectedOutput("Jrscc", "NKARe KJyiX", "✨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicDisk(t *testing.T) {
	writer := pine.NewWriter("ozFZB")
	writer.Disk("NsGRu %s", "ssVma")
	exp := generateExpectedOutput("ozFZB", "NsGRu ssVma", "💾", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraDisk(t *testing.T) {
	writer := pine.NewWriter("nwTKS")
	writer.DiskExtra("mVoiG", "sbOJi %s", "FQGZs")
	extraVal := "mVoiG"
	exp := generateExpectedOutput("nwTKS", "sbOJi FQGZs", "💾", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterDisk(t *testing.T) {
	writer := pine.NewWriter("updOM")
	ext := writer.WithExtra("eRVja")
	ext.Disk("LOpbU %s", "OpEdK")
	extraVal := "eRVja"
	exp := generateExpectedOutput("updOM", "LOpbU OpEdK", "💾", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicWTF(t *testing.T) {
	writer := pine.NewWriter("WKsXb")
	writer.WTF("RzLNT %s", "XYeUC")
	exp := generateExpectedOutput("WKsXb", "RzLNT XYeUC", "👻", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWTF(t *testing.T) {
	writer := pine.NewWriter("SJfjz")
	writer.WTFExtra("aLbtZ", "GyRAO %s", "mBTvK")
	extraVal := "aLbtZ"
	exp := generateExpectedOutput("SJfjz", "GyRAO mBTvK", "👻", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterWTF(t *testing.T) {
	writer := pine.NewWriter("QMDQi")
	ext := writer.WithExtra("YCOhg")
	ext.WTF("syMGe %s", "uDtRz")
	extraVal := "YCOhg"
	exp := generateExpectedOutput("QMDQi", "syMGe uDtRz", "👻", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSuccess(t *testing.T) {
	writer := pine.NewWriter("JHYNu")
	writer.Success("HOvgS %s", "eycJP")
	exp := generateExpectedOutput("JHYNu", "HOvgS eycJP", "✅", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSuccess(t *testing.T) {
	writer := pine.NewWriter("uSqfg")
	writer.SuccessExtra("qVMkP", "fNjJh %s", "hjUVR")
	extraVal := "qVMkP"
	exp := generateExpectedOutput("uSqfg", "fNjJh hjUVR", "✅", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSuccess(t *testing.T) {
	writer := pine.NewWriter("IZRgB")
	ext := writer.WithExtra("myArK")
	ext.Success("YVkUR %s", "UpiFv")
	extraVal := "myArK"
	exp := generateExpectedOutput("IZRgB", "YVkUR UpiFv", "✅", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicWarn(t *testing.T) {
	writer := pine.NewWriter("BjMkX")
	writer.Warn("Ctzkj %s", "kZIva")
	exp := generateExpectedOutput("BjMkX", "Ctzkj kZIva", "⚠️", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWarn(t *testing.T) {
	writer := pine.NewWriter("xyALB")
	writer.WarnExtra("sdjSG", "VbWGv %s", "bqzge")
	extraVal := "sdjSG"
	exp := generateExpectedOutput("xyALB", "VbWGv bqzge", "⚠️", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterWarn(t *testing.T) {
	writer := pine.NewWriter("IBuuf")
	ext := writer.WithExtra("FMoWd")
	ext.Warn("pngCw %s", "FkDif")
	extraVal := "FMoWd"
	exp := generateExpectedOutput("IBuuf", "pngCw FkDif", "⚠️", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicError(t *testing.T) {
	writer := pine.NewWriter("rTICT")
	writer.Error("iTskZ %s", "oQJMq")
	exp := generateExpectedOutput("rTICT", "iTskZ oQJMq", "🚨", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraError(t *testing.T) {
	writer := pine.NewWriter("yfroR")
	writer.ErrorExtra("ODMbN", "ojIYx %s", "yeSxZ")
	extraVal := "ODMbN"
	exp := generateExpectedOutput("yfroR", "ojIYx yeSxZ", "🚨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterError(t *testing.T) {
	writer := pine.NewWriter("PMHDt")
	ext := writer.WithExtra("JmHAY")
	ext.Error("DRZnP %s", "NRWCJ")
	extraVal := "JmHAY"
	exp := generateExpectedOutput("PMHDt", "DRZnP NRWCJ", "🚨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicTiming(t *testing.T) {
	writer := pine.NewWriter("VgzHb")
	writer.Timing("ORsUf %s", "UMAps")
	exp := generateExpectedOutput("VgzHb", "ORsUf UMAps", "⏱", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraTiming(t *testing.T) {
	writer := pine.NewWriter("fFbbG")
	writer.TimingExtra("Gcnqb", "lmYYt %s", "EjVgw")
	extraVal := "Gcnqb"
	exp := generateExpectedOutput("fFbbG", "lmYYt EjVgw", "⏱", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterTiming(t *testing.T) {
	writer := pine.NewWriter("XmZOt")
	ext := writer.WithExtra("aRLUt")
	ext.Timing("aEREu %s", "nUZjQ")
	extraVal := "aRLUt"
	exp := generateExpectedOutput("XmZOt", "aEREu nUZjQ", "⏱", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicInfo(t *testing.T) {
	writer := pine.NewWriter("DvoxI")
	writer.Info("mYgmS %s", "VYBAD")
	exp := generateExpectedOutput("DvoxI", "mYgmS VYBAD", "💬", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraInfo(t *testing.T) {
	writer := pine.NewWriter("IubeY")
	writer.InfoExtra("TNDtj", "fsfgP %s", "yCKmx")
	extraVal := "TNDtj"
	exp := generateExpectedOutput("IubeY", "fsfgP yCKmx", "💬", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterInfo(t *testing.T) {
	writer := pine.NewWriter("Lpruc")
	ext := writer.WithExtra("jiOgj")
	ext.Info("AyRRD %s", "edMiy")
	extraVal := "jiOgj"
	exp := generateExpectedOutput("Lpruc", "AyRRD edMiy", "💬", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}
