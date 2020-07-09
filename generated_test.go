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

func TestBasicBug(t *testing.T) {
	writer := pine.NewWriter("1f4304")
	writer.Bug("1f4104 %s", "1f4204")
	exp := generateExpectedOutput("1f4304", "1f4104 1f4204", "🐞", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraBug(t *testing.T) {
	writer := pine.NewWriter("1f4604")
	writer.BugExtra("1f4704", "1f4404 %s", "1f4504")
	extraVal := "1f4704"
	exp := generateExpectedOutput("1f4604", "1f4404 1f4504", "🐞", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterBug(t *testing.T) {
	writer := pine.NewWriter("1f4a04")
	ext := writer.WithExtra("244305")
	ext.Bug("1f4804 %s", "1f4904")
	extraVal := "244305"
	exp := generateExpectedOutput("1f4a04", "1f4804 1f4904", "🐞", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicDisk(t *testing.T) {
	writer := pine.NewWriter("1e8904")
	writer.Disk("1e8704 %s", "1e8804")
	exp := generateExpectedOutput("1e8904", "1e8704 1e8804", "📀", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraDisk(t *testing.T) {
	writer := pine.NewWriter("1e8c04")
	writer.DiskExtra("1e8d04", "1e8a04 %s", "1e8b04")
	extraVal := "1e8d04"
	exp := generateExpectedOutput("1e8c04", "1e8a04 1e8b04", "📀", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterDisk(t *testing.T) {
	writer := pine.NewWriter("1e9004")
	ext := writer.WithExtra("236e04")
	ext.Disk("1e8e04 %s", "1e8f04")
	extraVal := "236e04"
	exp := generateExpectedOutput("1e9004", "1e8e04 1e8f04", "📀", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicError(t *testing.T) {
	writer := pine.NewWriter("1f4004")
	writer.Error("1f3e04 %s", "1f3f04")
	exp := generateExpectedOutput("1f4004", "1f3e04 1f3f04", "🛑", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraError(t *testing.T) {
	writer := pine.NewWriter("1f4304")
	writer.ErrorExtra("1f4404", "1f4104 %s", "1f4204")
	extraVal := "1f4404"
	exp := generateExpectedOutput("1f4304", "1f4104 1f4204", "🛑", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterError(t *testing.T) {
	writer := pine.NewWriter("1f4704")
	ext := writer.WithExtra("243e04")
	ext.Error("1f4504 %s", "1f4604")
	extraVal := "243e04"
	exp := generateExpectedOutput("1f4704", "1f4504 1f4604", "🛑", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicFinish(t *testing.T) {
	writer := pine.NewWriter("1ec204")
	writer.Finish("1ec004 %s", "1ec104")
	exp := generateExpectedOutput("1ec204", "1ec004 1ec104", "🆗", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraFinish(t *testing.T) {
	writer := pine.NewWriter("1ec504")
	writer.FinishExtra("1ec604", "1ec304 %s", "1ec404")
	extraVal := "1ec604"
	exp := generateExpectedOutput("1ec504", "1ec304 1ec404", "🆗", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterFinish(t *testing.T) {
	writer := pine.NewWriter("1ec904")
	ext := writer.WithExtra("23b104")
	ext.Finish("1ec704 %s", "1ec804")
	extraVal := "23b104"
	exp := generateExpectedOutput("1ec904", "1ec704 1ec804", "🆗", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicInfo(t *testing.T) {
	writer := pine.NewWriter("1fb504")
	writer.Info("1fb304 %s", "1fb404")
	exp := generateExpectedOutput("1fb504", "1fb304 1fb404", "💬", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraInfo(t *testing.T) {
	writer := pine.NewWriter("1fb804")
	writer.InfoExtra("1fb904", "1fb604 %s", "1fb704")
	extraVal := "1fb904"
	exp := generateExpectedOutput("1fb804", "1fb604 1fb704", "💬", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterInfo(t *testing.T) {
	writer := pine.NewWriter("1fbc04")
	ext := writer.WithExtra("24c505")
	ext.Info("1fba04 %s", "1fbb04")
	extraVal := "24c505"
	exp := generateExpectedOutput("1fbc04", "1fba04 1fbb04", "💬", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicLock(t *testing.T) {
	writer := pine.NewWriter("1f0f04")
	writer.Lock("1f0d04 %s", "1f0e04")
	exp := generateExpectedOutput("1f0f04", "1f0d04 1f0e04", "🔒", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraLock(t *testing.T) {
	writer := pine.NewWriter("1f1204")
	writer.LockExtra("1f1304", "1f1004 %s", "1f1104")
	extraVal := "1f1304"
	exp := generateExpectedOutput("1f1204", "1f1004 1f1104", "🔒", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterLock(t *testing.T) {
	writer := pine.NewWriter("1f1604")
	ext := writer.WithExtra("240704")
	ext.Lock("1f1404 %s", "1f1504")
	extraVal := "240704"
	exp := generateExpectedOutput("1f1604", "1f1404 1f1504", "🔒", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicPlug(t *testing.T) {
	writer := pine.NewWriter("1ee504")
	writer.Plug("1ee304 %s", "1ee404")
	exp := generateExpectedOutput("1ee504", "1ee304 1ee404", "🔌", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraPlug(t *testing.T) {
	writer := pine.NewWriter("1ee804")
	writer.PlugExtra("1ee904", "1ee604 %s", "1ee704")
	extraVal := "1ee904"
	exp := generateExpectedOutput("1ee804", "1ee604 1ee704", "🔌", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterPlug(t *testing.T) {
	writer := pine.NewWriter("1eec04")
	ext := writer.WithExtra("23d704")
	ext.Plug("1eea04 %s", "1eeb04")
	extraVal := "23d704"
	exp := generateExpectedOutput("1eec04", "1eea04 1eeb04", "🔌", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSecure(t *testing.T) {
	writer := pine.NewWriter("1f0804")
	writer.Secure("1f0604 %s", "1f0704")
	exp := generateExpectedOutput("1f0804", "1f0604 1f0704", "🔑", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSecure(t *testing.T) {
	writer := pine.NewWriter("1f0b04")
	writer.SecureExtra("1f0c04", "1f0904 %s", "1f0a04")
	extraVal := "1f0c04"
	exp := generateExpectedOutput("1f0b04", "1f0904 1f0a04", "🔑", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSecure(t *testing.T) {
	writer := pine.NewWriter("1f0f04")
	ext := writer.WithExtra("23ff04")
	ext.Secure("1f0d04 %s", "1f0e04")
	extraVal := "23ff04"
	exp := generateExpectedOutput("1f0f04", "1f0d04 1f0e04", "🔑", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSleep(t *testing.T) {
	writer := pine.NewWriter("1f7d04")
	writer.Sleep("1f7b04 %s", "1f7c04")
	exp := generateExpectedOutput("1f7d04", "1f7b04 1f7c04", "💤", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSleep(t *testing.T) {
	writer := pine.NewWriter("1f8004")
	writer.SleepExtra("1f8104", "1f7e04 %s", "1f7f04")
	extraVal := "1f8104"
	exp := generateExpectedOutput("1f8004", "1f7e04 1f7f04", "💤", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSleep(t *testing.T) {
	writer := pine.NewWriter("1f8404")
	ext := writer.WithExtra("248505")
	ext.Sleep("1f8204 %s", "1f8304")
	extraVal := "248505"
	exp := generateExpectedOutput("1f8404", "1f8204 1f8304", "💤", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSpawn(t *testing.T) {
	writer := pine.NewWriter("18e304")
	writer.Spawn("18e104 %s", "18e204")
	exp := generateExpectedOutput("18e304", "18e104 18e204", "✨", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSpawn(t *testing.T) {
	writer := pine.NewWriter("18e604")
	writer.SpawnExtra("18e704", "18e404 %s", "18e504")
	extraVal := "18e704"
	exp := generateExpectedOutput("18e604", "18e404 18e504", "✨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSpawn(t *testing.T) {
	writer := pine.NewWriter("18ea04")
	ext := writer.WithExtra("1d4c04")
	ext.Spawn("18e804 %s", "18e904")
	extraVal := "1d4c04"
	exp := generateExpectedOutput("18ea04", "18e804 18e904", "✨", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicSuccess(t *testing.T) {
	writer := pine.NewWriter("1fd704")
	writer.Success("1fd504 %s", "1fd604")
	exp := generateExpectedOutput("1fd704", "1fd504 1fd604", "🟢", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraSuccess(t *testing.T) {
	writer := pine.NewWriter("1fda04")
	writer.SuccessExtra("1fdb04", "1fd804 %s", "1fd904")
	extraVal := "1fdb04"
	exp := generateExpectedOutput("1fda04", "1fd804 1fd904", "🟢", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterSuccess(t *testing.T) {
	writer := pine.NewWriter("1fde04")
	ext := writer.WithExtra("24ea05")
	ext.Success("1fdc04 %s", "1fdd04")
	extraVal := "24ea05"
	exp := generateExpectedOutput("1fde04", "1fdc04 1fdd04", "🟢", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicTerminate(t *testing.T) {
	writer := pine.NewWriter("18e604")
	writer.Terminate("18e404 %s", "18e504")
	exp := generateExpectedOutput("18e604", "18e404 18e504", "⭕", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraTerminate(t *testing.T) {
	writer := pine.NewWriter("18e904")
	writer.TerminateExtra("18ea04", "18e704 %s", "18e804")
	extraVal := "18ea04"
	exp := generateExpectedOutput("18e904", "18e704 18e804", "⭕", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterTerminate(t *testing.T) {
	writer := pine.NewWriter("18ed04")
	ext := writer.WithExtra("1d4d04")
	ext.Terminate("18eb04 %s", "18ec04")
	extraVal := "1d4d04"
	exp := generateExpectedOutput("18ed04", "18eb04 18ec04", "⭕", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicThread(t *testing.T) {
	writer := pine.NewWriter("209c05")
	writer.Thread("209a04 %s", "209b04")
	exp := generateExpectedOutput("209c05", "209a04 209b04", "🧵", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraThread(t *testing.T) {
	writer := pine.NewWriter("209f05")
	writer.ThreadExtra("20a005", "209d05 %s", "209e05")
	extraVal := "20a005"
	exp := generateExpectedOutput("209f05", "209d05 209e05", "🧵", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterThread(t *testing.T) {
	writer := pine.NewWriter("20a305")
	ext := writer.WithExtra("25ca05")
	ext.Thread("20a105 %s", "20a205")
	extraVal := "25ca05"
	exp := generateExpectedOutput("20a305", "20a105 20a205", "🧵", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicTiming(t *testing.T) {
	writer := pine.NewWriter("18ba04")
	writer.Timing("18b804 %s", "18b904")
	exp := generateExpectedOutput("18ba04", "18b804 18b904", "⏱", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraTiming(t *testing.T) {
	writer := pine.NewWriter("18bd04")
	writer.TimingExtra("18be04", "18bb04 %s", "18bc04")
	extraVal := "18be04"
	exp := generateExpectedOutput("18bd04", "18bb04 18bc04", "⏱", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterTiming(t *testing.T) {
	writer := pine.NewWriter("18c104")
	ext := writer.WithExtra("1d1f04")
	ext.Timing("18bf04 %s", "18c004")
	extraVal := "1d1f04"
	exp := generateExpectedOutput("18c104", "18bf04 18c004", "⏱", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicWTF(t *testing.T) {
	writer := pine.NewWriter("201604")
	writer.WTF("201404 %s", "201504")
	exp := generateExpectedOutput("201604", "201404 201504", "👻", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWTF(t *testing.T) {
	writer := pine.NewWriter("201904")
	writer.WTFExtra("201a04", "201704 %s", "201804")
	extraVal := "201a04"
	exp := generateExpectedOutput("201904", "201704 201804", "👻", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterWTF(t *testing.T) {
	writer := pine.NewWriter("201d04")
	ext := writer.WithExtra("253405")
	ext.WTF("201b04 %s", "201c04")
	extraVal := "253405"
	exp := generateExpectedOutput("201d04", "201b04 201c04", "👻", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestBasicWarn(t *testing.T) {
	writer := pine.NewWriter("1fd004")
	writer.Warn("1fce04 %s", "1fcf04")
	exp := generateExpectedOutput("1fd004", "1fce04 1fcf04", "🟡", nil)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWarn(t *testing.T) {
	writer := pine.NewWriter("1fd304")
	writer.WarnExtra("1fd404", "1fd104 %s", "1fd204")
	extraVal := "1fd404"
	exp := generateExpectedOutput("1fd304", "1fd104 1fd204", "🟡", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}

func TestExtraWriterWarn(t *testing.T) {
	writer := pine.NewWriter("1fd704")
	ext := writer.WithExtra("24e205")
	ext.Warn("1fd504 %s", "1fd604")
	extraVal := "24e205"
	exp := generateExpectedOutput("1fd704", "1fd504 1fd604", "🟡", &extraVal)
	if lastOutput != exp {
		t.Errorf("Failed. Expectation was %s, got instead %s", exp, lastOutput)
	}
}
