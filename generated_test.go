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

func TestBasicError(t *testing.T) {
	writer := pine.NewWriter("MRAjW")
	writer.Error("XVlBz %s", "gbaiC")
	exp := generateExpectedOutput("MRAjW", "XVlBz gbaiC", "🚨", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraError(t *testing.T) {
	writer := pine.NewWriter("hxKQF")
	writer.ErrorExtra("DaFpL", "whTHc %s", "tcuAx")
	extraVal := "DaFpL"
	exp := generateExpectedOutput("hxKQF", "whTHc tcuAx", "🚨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterError(t *testing.T) {
	writer := pine.NewWriter("RsWxP")
	ext := writer.WithExtra("LDnJO")
	ext.Error("SjFbc %s", "XoEFf")
	extraVal := "LDnJO"
	exp := generateExpectedOutput("RsWxP", "SjFbc XoEFf", "🚨", &extraVal)
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

func TestBasicDisk(t *testing.T) {
	writer := pine.NewWriter("EkXBA")
	writer.Disk("krBEm %s", "fdzdc")
	exp := generateExpectedOutput("EkXBA", "krBEm fdzdc", "💾", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraDisk(t *testing.T) {
	writer := pine.NewWriter("TCoaN")
	writer.DiskExtra("atyyi", "kjQZL %s", "CtTMt")
	extraVal := "atyyi"
	exp := generateExpectedOutput("TCoaN", "kjQZL CtTMt", "💾", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterDisk(t *testing.T) {
	writer := pine.NewWriter("Jrscc")
	ext := writer.WithExtra("tNswY")
	ext.Disk("NKARe %s", "KJyiX")
	extraVal := "tNswY"
	exp := generateExpectedOutput("Jrscc", "NKARe KJyiX", "💾", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicWTF(t *testing.T) {
	writer := pine.NewWriter("ozFZB")
	writer.WTF("NsGRu %s", "ssVma")
	exp := generateExpectedOutput("ozFZB", "NsGRu ssVma", "👻", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWTF(t *testing.T) {
	writer := pine.NewWriter("nwTKS")
	writer.WTFExtra("mVoiG", "sbOJi %s", "FQGZs")
	extraVal := "mVoiG"
	exp := generateExpectedOutput("nwTKS", "sbOJi FQGZs", "👻", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterWTF(t *testing.T) {
	writer := pine.NewWriter("updOM")
	ext := writer.WithExtra("eRVja")
	ext.WTF("LOpbU %s", "OpEdK")
	extraVal := "eRVja"
	exp := generateExpectedOutput("updOM", "LOpbU OpEdK", "👻", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicLock(t *testing.T) {
	writer := pine.NewWriter("WKsXb")
	writer.Lock("RzLNT %s", "XYeUC")
	exp := generateExpectedOutput("WKsXb", "RzLNT XYeUC", "🔒", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraLock(t *testing.T) {
	writer := pine.NewWriter("SJfjz")
	writer.LockExtra("aLbtZ", "GyRAO %s", "mBTvK")
	extraVal := "aLbtZ"
	exp := generateExpectedOutput("SJfjz", "GyRAO mBTvK", "🔒", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterLock(t *testing.T) {
	writer := pine.NewWriter("QMDQi")
	ext := writer.WithExtra("YCOhg")
	ext.Lock("syMGe %s", "uDtRz")
	extraVal := "YCOhg"
	exp := generateExpectedOutput("QMDQi", "syMGe uDtRz", "🔒", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSecure(t *testing.T) {
	writer := pine.NewWriter("JHYNu")
	writer.Secure("HOvgS %s", "eycJP")
	exp := generateExpectedOutput("JHYNu", "HOvgS eycJP", "🔑", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSecure(t *testing.T) {
	writer := pine.NewWriter("uSqfg")
	writer.SecureExtra("qVMkP", "fNjJh %s", "hjUVR")
	extraVal := "qVMkP"
	exp := generateExpectedOutput("uSqfg", "fNjJh hjUVR", "🔑", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSecure(t *testing.T) {
	writer := pine.NewWriter("IZRgB")
	ext := writer.WithExtra("myArK")
	ext.Secure("YVkUR %s", "UpiFv")
	extraVal := "myArK"
	exp := generateExpectedOutput("IZRgB", "YVkUR UpiFv", "🔑", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSuccess(t *testing.T) {
	writer := pine.NewWriter("BjMkX")
	writer.Success("Ctzkj %s", "kZIva")
	exp := generateExpectedOutput("BjMkX", "Ctzkj kZIva", "✅", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSuccess(t *testing.T) {
	writer := pine.NewWriter("xyALB")
	writer.SuccessExtra("sdjSG", "VbWGv %s", "bqzge")
	extraVal := "sdjSG"
	exp := generateExpectedOutput("xyALB", "VbWGv bqzge", "✅", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSuccess(t *testing.T) {
	writer := pine.NewWriter("IBuuf")
	ext := writer.WithExtra("FMoWd")
	ext.Success("pngCw %s", "FkDif")
	extraVal := "FMoWd"
	exp := generateExpectedOutput("IBuuf", "pngCw FkDif", "✅", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicWarn(t *testing.T) {
	writer := pine.NewWriter("rTICT")
	writer.Warn("iTskZ %s", "oQJMq")
	exp := generateExpectedOutput("rTICT", "iTskZ oQJMq", "⚠️", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWarn(t *testing.T) {
	writer := pine.NewWriter("yfroR")
	writer.WarnExtra("ODMbN", "ojIYx %s", "yeSxZ")
	extraVal := "ODMbN"
	exp := generateExpectedOutput("yfroR", "ojIYx yeSxZ", "⚠️", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterWarn(t *testing.T) {
	writer := pine.NewWriter("PMHDt")
	ext := writer.WithExtra("JmHAY")
	ext.Warn("DRZnP %s", "NRWCJ")
	extraVal := "JmHAY"
	exp := generateExpectedOutput("PMHDt", "DRZnP NRWCJ", "⚠️", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicFinish(t *testing.T) {
	writer := pine.NewWriter("VgzHb")
	writer.Finish("ORsUf %s", "UMAps")
	exp := generateExpectedOutput("VgzHb", "ORsUf UMAps", "🏁", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraFinish(t *testing.T) {
	writer := pine.NewWriter("fFbbG")
	writer.FinishExtra("Gcnqb", "lmYYt %s", "EjVgw")
	extraVal := "Gcnqb"
	exp := generateExpectedOutput("fFbbG", "lmYYt EjVgw", "🏁", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterFinish(t *testing.T) {
	writer := pine.NewWriter("XmZOt")
	ext := writer.WithExtra("aRLUt")
	ext.Finish("aEREu %s", "nUZjQ")
	extraVal := "aRLUt"
	exp := generateExpectedOutput("XmZOt", "aEREu nUZjQ", "🏁", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSpawn(t *testing.T) {
	writer := pine.NewWriter("DvoxI")
	writer.Spawn("mYgmS %s", "VYBAD")
	exp := generateExpectedOutput("DvoxI", "mYgmS VYBAD", "✨", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSpawn(t *testing.T) {
	writer := pine.NewWriter("IubeY")
	writer.SpawnExtra("TNDtj", "fsfgP %s", "yCKmx")
	extraVal := "TNDtj"
	exp := generateExpectedOutput("IubeY", "fsfgP yCKmx", "✨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSpawn(t *testing.T) {
	writer := pine.NewWriter("Lpruc")
	ext := writer.WithExtra("jiOgj")
	ext.Spawn("AyRRD %s", "edMiy")
	extraVal := "jiOgj"
	exp := generateExpectedOutput("Lpruc", "AyRRD edMiy", "✨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicTiming(t *testing.T) {
	writer := pine.NewWriter("frDGX")
	writer.Timing("hYeVw %s", "BTCML")
	exp := generateExpectedOutput("frDGX", "hYeVw BTCML", "⏱", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraTiming(t *testing.T) {
	writer := pine.NewWriter("cLVCx")
	writer.TimingExtra("aSJlD", "qwpzw %s", "VGqMZ")
	extraVal := "aSJlD"
	exp := generateExpectedOutput("cLVCx", "qwpzw VGqMZ", "⏱", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterTiming(t *testing.T) {
	writer := pine.NewWriter("qkKHq")
	ext := writer.WithExtra("gBpnb")
	ext.Timing("SYEof %s", "kkEYe")
	extraVal := "gBpnb"
	exp := generateExpectedOutput("qkKHq", "SYEof kkEYe", "⏱", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicInfo(t *testing.T) {
	writer := pine.NewWriter("UMmpB")
	writer.Info("PbgHM %s", "LUIDj")
	exp := generateExpectedOutput("UMmpB", "PbgHM LUIDj", "💬", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraInfo(t *testing.T) {
	writer := pine.NewWriter("aiIsN")
	writer.InfoExtra("BakqS", "HCSjM %s", "Jjxzu")
	extraVal := "BakqS"
	exp := generateExpectedOutput("aiIsN", "HCSjM Jjxzu", "💬", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterInfo(t *testing.T) {
	writer := pine.NewWriter("aczAI")
	ext := writer.WithExtra("nLqLI")
	ext.Info("wQpOQ %s", "gNczg")
	extraVal := "nLqLI"
	exp := generateExpectedOutput("aczAI", "wQpOQ gNczg", "💬", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}
