;================================================
; WSL Path Manager - NSIS Installer Script
;================================================
; Standard Windows installer for WSL Path Manager
; Built with NSIS (Nullsoft Scriptable Install System)
;
; Build command:
;   wails build --target windows/amd64 --nsis
;
; For manual build with pre-compiled binary:
;   makensis -DARG_WAILS_AMD64_BINARY=..\..\bin\wslpathmanager.exe project.nsi
;================================================

Unicode true

;------------------------------------------------
; Overridable Project Info
; (wails_tools.nsh populates these if not defined)
;------------------------------------------------
!define INFO_PROJECTNAME    "wslpathmanager"
!define INFO_COMPANYNAME    "NextXierra"
!define INFO_PRODUCTNAME    "WSL Path Manager"
!define INFO_PRODUCTVERSION "0.0.2"
!define INFO_COPYRIGHT      "Copyright (c) 2025 NextXierra"

; Installer metadata
!define PRODUCT_EXECUTABLE  "wslpathmanager.exe"
!define UNINST_KEY_NAME    "NextXierraWSLPathManager"
!define REQUEST_EXECUTION_LEVEL "admin"

;------------------------------------------------
; Include Wails NSIS tools
;------------------------------------------------
!include "wails_tools.nsh"

;------------------------------------------------
; Version Info (4-part version required by NSIS)
;------------------------------------------------
VIProductVersion "${INFO_PRODUCTVERSION}.0"
VIFileVersion    "${INFO_PRODUCTVERSION}.0"

VIAddVersionKey "CompanyName"     "${INFO_COMPANYNAME}"
VIAddVersionKey "FileDescription" "${INFO_PRODUCTNAME} Installer"
VIAddVersionKey "ProductVersion"  "${INFO_PRODUCTVERSION}"
VIAddVersionKey "FileVersion"     "${INFO_PRODUCTVERSION}"
VIAddVersionKey "LegalCopyright"  "${INFO_COPYRIGHT}"
VIAddVersionKey "ProductName"     "${INFO_PRODUCTNAME}"

;------------------------------------------------
; Installer Attributes
;------------------------------------------------
Name "${INFO_PRODUCTNAME}"
OutFile "..\..\bin\${INFO_PROJECTNAME}-${ARCH}-setup.exe"
InstallDir "$PROGRAMFILES64\${INFO_COMPANYNAME}\${INFO_PRODUCTNAME}"
InstallDirRegKey HKLM "Software\${INFO_COMPANYNAME}\${INFO_PRODUCTNAME}" "InstallPath"

ShowInstDetails show
ShowUnInstDetails show

;------------------------------------------------
; Modern UI (MUI2)
;------------------------------------------------
!include "MUI2.nsh"

; Installer icon (uses the app's icon)
!define MUI_ICON "..\icon.ico"
!define MUI_UNICON "..\icon.ico"

; Installer pages
!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_LICENSE "..\..\LICENSE.txt"
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH

; Uninstaller pages
!insertmacro MUI_UNPAGE_CONFIRM
!insertmacro MUI_UNPAGE_INSTFILES

; Language
!insertmacro MUI_LANGUAGE "English"

;------------------------------------------------
; Installer Sections
;------------------------------------------------
Section "Install" SecMain
    !insertmacro wails.setShellContext
    !insertmacro wails.webview2runtime

    SetOutPath "$INSTDIR"

    ; Install main executable
    !insertmacro wails.files

    ; Install supporting files (if any in extra files)
    ; File /nonfatal "README.txt"
    ; File "..\README.txt"

    ; Create Start Menu shortcuts
    CreateDirectory "$SMPROGRAMS\${INFO_PRODUCTNAME}"
    CreateShortcut "$SMPROGRAMS\${INFO_PRODUCTNAME}\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"
    CreateShortcut "$SMPROGRAMS\${INFO_PRODUCTNAME}\Uninstall.lnk" "$INSTDIR\uninstall.exe"

    ; Create Desktop shortcut
    CreateShortcut "$DESKTOP\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\${PRODUCT_EXECUTABLE}"

    ; Write registry info for Add/Remove Programs
    WriteRegStr HKLM "Software\${INFO_COMPANYNAME}\${INFO_PRODUCTNAME}" "InstallPath" "$INSTDIR"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "Publisher" "${INFO_COMPANYNAME}"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "DisplayName" "${INFO_PRODUCTNAME}"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "DisplayVersion" "${INFO_PRODUCTVERSION}"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "DisplayIcon" "$INSTDIR\${PRODUCT_EXECUTABLE}"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "UninstallString" '"$INSTDIR\uninstall.exe"'
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "QuietUninstallString" '"$INSTDIR\uninstall.exe" /S'
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "InstallLocation" "$INSTDIR"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "URLInfoAbout" "https://github.com/NextXierra/wslpathmanager"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "HelpLink" "https://github.com/NextXierra/wslpathmanager/issues"

    ; Calculate installed size
    ${GetSize} "$INSTDIR" "/S=0K" $0 $1 $2
    IntFmt $0 "0x%08X" $0
    WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}" "EstimatedSize" "$0"

    ; Write uninstaller
    !insertmacro wails.writeUninstaller

    ; File associations (if any defined in wails.json)
    !insertmacro wails.associateFiles
    !insertmacro wails.associateCustomProtocols
SectionEnd

;------------------------------------------------
; Uninstaller Section
;------------------------------------------------
Section "uninstall"
    !insertmacro wails.setShellContext

    ; Remove WebView2 data
    RMDir /r "$AppData\${PRODUCT_EXECUTABLE}"

    ; Remove installed files
    RMDir /r "$INSTDIR"

    ; Remove Start Menu shortcuts
    RMDir /r "$SMPROGRAMS\${INFO_PRODUCTNAME}"

    ; Remove Desktop shortcut
    Delete "$DESKTOP\${INFO_PRODUCTNAME}.lnk"

    ; Remove registry entries
    DeleteRegKey HKLM "Software\${INFO_COMPANYNAME}\${INFO_PRODUCTNAME}"
    DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${UNINST_KEY_NAME}"

    ; Remove file associations
    !insertmacro wails.unassociateFiles
    !insertmacro wails.unassociateCustomProtocols

    ; Remove uninstaller
    !insertmacro wails.deleteUninstaller
SectionEnd

;------------------------------------------------
; Initialization
;------------------------------------------------
Function .onInit
    !insertmacro wails.checkArchitecture
FunctionEnd
