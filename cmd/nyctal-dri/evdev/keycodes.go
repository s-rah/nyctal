// Linux evdev event codes (keys and buttons)
package evdev

import "fmt"

// key/button code
type KeyOrButton uint16

func (k KeyOrButton) String() string {
	str, found := keyToString[k]
	if !found {
		return fmt.Sprintf("Unknown(%d)", k)
	}

	return str
}

// all taken from https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
// (defines with "KEY_" or "BTN_" prefix)

// https://github.com/torvalds/linux/blob/5c8fe583cc/include/uapi/linux/input-event-codes.h#L65

const (
	KeyRESERVED                 KeyOrButton = 0
	KeyESC                      KeyOrButton = 1
	Key1                        KeyOrButton = 2
	Key2                        KeyOrButton = 3
	Key3                        KeyOrButton = 4
	Key4                        KeyOrButton = 5
	Key5                        KeyOrButton = 6
	Key6                        KeyOrButton = 7
	Key7                        KeyOrButton = 8
	Key8                        KeyOrButton = 9
	Key9                        KeyOrButton = 10
	Key0                        KeyOrButton = 11
	KeyMINUS                    KeyOrButton = 12
	KeyEQUAL                    KeyOrButton = 13
	KeyBACKSPACE                KeyOrButton = 14
	KeyTAB                      KeyOrButton = 15
	KeyQ                        KeyOrButton = 16
	KeyW                        KeyOrButton = 17
	KeyE                        KeyOrButton = 18
	KeyR                        KeyOrButton = 19
	KeyT                        KeyOrButton = 20
	KeyY                        KeyOrButton = 21
	KeyU                        KeyOrButton = 22
	KeyI                        KeyOrButton = 23
	KeyO                        KeyOrButton = 24
	KeyP                        KeyOrButton = 25
	KeyLEFTBRACE                KeyOrButton = 26
	KeyRIGHTBRACE               KeyOrButton = 27
	KeyENTER                    KeyOrButton = 28
	KeyLEFTCTRL                 KeyOrButton = 29
	KeyA                        KeyOrButton = 30
	KeyS                        KeyOrButton = 31
	KeyD                        KeyOrButton = 32
	KeyF                        KeyOrButton = 33
	KeyG                        KeyOrButton = 34
	KeyH                        KeyOrButton = 35
	KeyJ                        KeyOrButton = 36
	KeyK                        KeyOrButton = 37
	KeyL                        KeyOrButton = 38
	KeySEMICOLON                KeyOrButton = 39
	KeyAPOSTROPHE               KeyOrButton = 40
	KeyGRAVE                    KeyOrButton = 41
	KeyLEFTSHIFT                KeyOrButton = 42
	KeyBACKSLASH                KeyOrButton = 43
	KeyZ                        KeyOrButton = 44
	KeyX                        KeyOrButton = 45
	KeyC                        KeyOrButton = 46
	KeyV                        KeyOrButton = 47
	KeyB                        KeyOrButton = 48
	KeyN                        KeyOrButton = 49
	KeyM                        KeyOrButton = 50
	KeyCOMMA                    KeyOrButton = 51
	KeyDOT                      KeyOrButton = 52
	KeySLASH                    KeyOrButton = 53
	KeyRIGHTSHIFT               KeyOrButton = 54
	KeyKPASTERISK               KeyOrButton = 55
	KeyLEFTALT                  KeyOrButton = 56
	KeySPACE                    KeyOrButton = 57
	KeyCAPSLOCK                 KeyOrButton = 58
	KeyF1                       KeyOrButton = 59
	KeyF2                       KeyOrButton = 60
	KeyF3                       KeyOrButton = 61
	KeyF4                       KeyOrButton = 62
	KeyF5                       KeyOrButton = 63
	KeyF6                       KeyOrButton = 64
	KeyF7                       KeyOrButton = 65
	KeyF8                       KeyOrButton = 66
	KeyF9                       KeyOrButton = 67
	KeyF10                      KeyOrButton = 68
	KeyNUMLOCK                  KeyOrButton = 69
	KeySCROLLLOCK               KeyOrButton = 70
	KeyKP7                      KeyOrButton = 71
	KeyKP8                      KeyOrButton = 72
	KeyKP9                      KeyOrButton = 73
	KeyKPMINUS                  KeyOrButton = 74
	KeyKP4                      KeyOrButton = 75
	KeyKP5                      KeyOrButton = 76
	KeyKP6                      KeyOrButton = 77
	KeyKPPLUS                   KeyOrButton = 78
	KeyKP1                      KeyOrButton = 79
	KeyKP2                      KeyOrButton = 80
	KeyKP3                      KeyOrButton = 81
	KeyKP0                      KeyOrButton = 82
	KeyKPDOT                    KeyOrButton = 83
	KeyZENKAKUHANKAKU           KeyOrButton = 85
	Key102ND                    KeyOrButton = 86
	KeyF11                      KeyOrButton = 87
	KeyF12                      KeyOrButton = 88
	KeyRO                       KeyOrButton = 89
	KeyKATAKANA                 KeyOrButton = 90
	KeyHIRAGANA                 KeyOrButton = 91
	KeyHENKAN                   KeyOrButton = 92
	KeyKATAKANAHIRAGANA         KeyOrButton = 93
	KeyMUHENKAN                 KeyOrButton = 94
	KeyKPJPCOMMA                KeyOrButton = 95
	KeyKPENTER                  KeyOrButton = 96
	KeyRIGHTCTRL                KeyOrButton = 97
	KeyKPSLASH                  KeyOrButton = 98
	KeySYSRQ                    KeyOrButton = 99
	KeyRIGHTALT                 KeyOrButton = 100
	KeyLINEFEED                 KeyOrButton = 101
	KeyHOME                     KeyOrButton = 102
	KeyUP                       KeyOrButton = 103
	KeyPAGEUP                   KeyOrButton = 104
	KeyLEFT                     KeyOrButton = 105
	KeyRIGHT                    KeyOrButton = 106
	KeyEND                      KeyOrButton = 107
	KeyDOWN                     KeyOrButton = 108
	KeyPAGEDOWN                 KeyOrButton = 109
	KeyINSERT                   KeyOrButton = 110
	KeyDELETE                   KeyOrButton = 111
	KeyMACRO                    KeyOrButton = 112
	KeyMUTE                     KeyOrButton = 113
	KeyVOLUMEDOWN               KeyOrButton = 114
	KeyVOLUMEUP                 KeyOrButton = 115
	KeyPOWER                    KeyOrButton = 116 // SC System Power Down
	KeyKPEQUAL                  KeyOrButton = 117
	KeyKPPLUSMINUS              KeyOrButton = 118
	KeyPAUSE                    KeyOrButton = 119
	KeySCALE                    KeyOrButton = 120 // AL Compiz Scale (Expose)
	KeyKPCOMMA                  KeyOrButton = 121
	KeyHANGEUL                  KeyOrButton = 122
	KeyHANGUEL                  KeyOrButton = KeyHANGEUL
	KeyHANJA                    KeyOrButton = 123
	KeyYEN                      KeyOrButton = 124
	KeyLEFTMETA                 KeyOrButton = 125
	KeyRIGHTMETA                KeyOrButton = 126
	KeyCOMPOSE                  KeyOrButton = 127
	KeySTOP                     KeyOrButton = 128 // AC Stop
	KeyAGAIN                    KeyOrButton = 129
	KeyPROPS                    KeyOrButton = 130 // AC Properties
	KeyUNDO                     KeyOrButton = 131 // AC Undo
	KeyFRONT                    KeyOrButton = 132
	KeyCOPY                     KeyOrButton = 133 // AC Copy
	KeyOPEN                     KeyOrButton = 134 // AC Open
	KeyPASTE                    KeyOrButton = 135 // AC Paste
	KeyFIND                     KeyOrButton = 136 // AC Search
	KeyCUT                      KeyOrButton = 137 // AC Cut
	KeyHELP                     KeyOrButton = 138 // AL Integrated Help Center
	KeyMENU                     KeyOrButton = 139 // Menu (show menu)
	KeyCALC                     KeyOrButton = 140 // AL Calculator
	KeySETUP                    KeyOrButton = 141
	KeySLEEP                    KeyOrButton = 142 // SC System Sleep
	KeyWAKEUP                   KeyOrButton = 143 // System Wake Up
	KeyFILE                     KeyOrButton = 144 // AL Local Machine Browser
	KeySENDFILE                 KeyOrButton = 145
	KeyDELETEFILE               KeyOrButton = 146
	KeyXFER                     KeyOrButton = 147
	KeyPROG1                    KeyOrButton = 148
	KeyPROG2                    KeyOrButton = 149
	KeyWWW                      KeyOrButton = 150 // AL Internet Browser
	KeyMSDOS                    KeyOrButton = 151
	KeyCOFFEE                   KeyOrButton = 152 // AL Terminal Lock/Screensaver
	KeySCREENLOCK               KeyOrButton = KeyCOFFEE
	KeyROTATE_DISPLAY           KeyOrButton = 153 // Display orientation for e.g. tablets
	KeyDIRECTION                KeyOrButton = KeyROTATE_DISPLAY
	KeyCYCLEWINDOWS             KeyOrButton = 154
	KeyMAIL                     KeyOrButton = 155
	KeyBOOKMARKS                KeyOrButton = 156 // AC Bookmarks
	KeyCOMPUTER                 KeyOrButton = 157
	KeyBACK                     KeyOrButton = 158 // AC Back
	KeyFORWARD                  KeyOrButton = 159 // AC Forward
	KeyCLOSECD                  KeyOrButton = 160
	KeyEJECTCD                  KeyOrButton = 161
	KeyEJECTCLOSECD             KeyOrButton = 162
	KeyNEXTSONG                 KeyOrButton = 163
	KeyPLAYPAUSE                KeyOrButton = 164
	KeyPREVIOUSSONG             KeyOrButton = 165
	KeySTOPCD                   KeyOrButton = 166
	KeyRECORD                   KeyOrButton = 167
	KeyREWIND                   KeyOrButton = 168
	KeyPHONE                    KeyOrButton = 169 // Media Select Telephone
	KeyISO                      KeyOrButton = 170
	KeyCONFIG                   KeyOrButton = 171 // AL Consumer Control Configuration
	KeyHOMEPAGE                 KeyOrButton = 172 // AC Home
	KeyREFRESH                  KeyOrButton = 173 // AC Refresh
	KeyEXIT                     KeyOrButton = 174 // AC Exit
	KeyMOVE                     KeyOrButton = 175
	KeyEDIT                     KeyOrButton = 176
	KeySCROLLUP                 KeyOrButton = 177
	KeySCROLLDOWN               KeyOrButton = 178
	KeyKPLEFTPAREN              KeyOrButton = 179
	KeyKPRIGHTPAREN             KeyOrButton = 180
	KeyNEW                      KeyOrButton = 181 // AC New
	KeyREDO                     KeyOrButton = 182 // AC Redo/Repeat
	KeyF13                      KeyOrButton = 183
	KeyF14                      KeyOrButton = 184
	KeyF15                      KeyOrButton = 185
	KeyF16                      KeyOrButton = 186
	KeyF17                      KeyOrButton = 187
	KeyF18                      KeyOrButton = 188
	KeyF19                      KeyOrButton = 189
	KeyF20                      KeyOrButton = 190
	KeyF21                      KeyOrButton = 191
	KeyF22                      KeyOrButton = 192
	KeyF23                      KeyOrButton = 193
	KeyF24                      KeyOrButton = 194
	KeyPLAYCD                   KeyOrButton = 200
	KeyPAUSECD                  KeyOrButton = 201
	KeyPROG3                    KeyOrButton = 202
	KeyPROG4                    KeyOrButton = 203
	KeyDASHBOARD                KeyOrButton = 204 // AL Dashboard
	KeySUSPEND                  KeyOrButton = 205
	KeyCLOSE                    KeyOrButton = 206 // AC Close
	KeyPLAY                     KeyOrButton = 207
	KeyFASTFORWARD              KeyOrButton = 208
	KeyBASSBOOST                KeyOrButton = 209
	KeyPRINT                    KeyOrButton = 210 // AC Print
	KeyHP                       KeyOrButton = 211
	KeyCAMERA                   KeyOrButton = 212
	KeySOUND                    KeyOrButton = 213
	KeyQUESTION                 KeyOrButton = 214
	KeyEMAIL                    KeyOrButton = 215
	KeyCHAT                     KeyOrButton = 216
	KeySEARCH                   KeyOrButton = 217
	KeyCONNECT                  KeyOrButton = 218
	KeyFINANCE                  KeyOrButton = 219 // AL Checkbook/Finance
	KeySPORT                    KeyOrButton = 220
	KeySHOP                     KeyOrButton = 221
	KeyALTERASE                 KeyOrButton = 222
	KeyCANCEL                   KeyOrButton = 223 // AC Cancel
	KeyBRIGHTNESSDOWN           KeyOrButton = 224
	KeyBRIGHTNESSUP             KeyOrButton = 225
	KeyMEDIA                    KeyOrButton = 226
	KeySWITCHVIDEOMODE          KeyOrButton = 227 // Cycle between available vi
	KeyKBDILLUMTOGGLE           KeyOrButton = 228
	KeyKBDILLUMDOWN             KeyOrButton = 229
	KeyKBDILLUMUP               KeyOrButton = 230
	KeySEND                     KeyOrButton = 231 // AC Send
	KeyREPLY                    KeyOrButton = 232 // AC Reply
	KeyFORWARDMAIL              KeyOrButton = 233 // AC Forward Msg
	KeySAVE                     KeyOrButton = 234 // AC Save
	KeyDOCUMENTS                KeyOrButton = 235
	KeyBATTERY                  KeyOrButton = 236
	KeyBLUETOOTH                KeyOrButton = 237
	KeyWLAN                     KeyOrButton = 238
	KeyUWB                      KeyOrButton = 239
	KeyUNKNOWN                  KeyOrButton = 240
	KeyVIDEO_NEXT               KeyOrButton = 241 // drive next video source
	KeyVIDEO_PREV               KeyOrButton = 242 // drive previous video source
	KeyBRIGHTNESS_CYCLE         KeyOrButton = 243 // brightness up, after max is min
	KeyBRIGHTNESS_AUTO          KeyOrButton = 244 // Set Auto Brightness: man
	KeyBRIGHTNESS_ZERO          KeyOrButton = KeyBRIGHTNESS_AUTO
	KeyDISPLAY_OFF              KeyOrButton = 245 // display device to off state
	KeyWWAN                     KeyOrButton = 246 // Wireless WAN (LTE, UMTS, GSM, etc.)
	KeyWIMAX                    KeyOrButton = KeyWWAN
	KeyRFKILL                   KeyOrButton = 247 // Key that controls all radios
	KeyMICMUTE                  KeyOrButton = 248 // Mute / unmute the microphone
	KeyOK                       KeyOrButton = 0x160
	KeySELECT                   KeyOrButton = 0x161
	KeyGOTO                     KeyOrButton = 0x162
	KeyCLEAR                    KeyOrButton = 0x163
	KeyPOWER2                   KeyOrButton = 0x164
	KeyOPTION                   KeyOrButton = 0x165
	KeyINFO                     KeyOrButton = 0x166 // AL OEM Features/Tips/Tutorial
	KeyTIME                     KeyOrButton = 0x167
	KeyVENDOR                   KeyOrButton = 0x168
	KeyARCHIVE                  KeyOrButton = 0x169
	KeyPROGRAM                  KeyOrButton = 0x16a // Media Select Program Guide
	KeyCHANNEL                  KeyOrButton = 0x16b
	KeyFAVORITES                KeyOrButton = 0x16c
	KeyEPG                      KeyOrButton = 0x16d
	KeyPVR                      KeyOrButton = 0x16e // Media Select Home
	KeyMHP                      KeyOrButton = 0x16f
	KeyLANGUAGE                 KeyOrButton = 0x170
	KeyTITLE                    KeyOrButton = 0x171
	KeySUBTITLE                 KeyOrButton = 0x172
	KeyANGLE                    KeyOrButton = 0x173
	KeyFULL_SCREEN              KeyOrButton = 0x174 // AC View Toggle
	KeyZOOM                     KeyOrButton = KeyFULL_SCREEN
	KeyMODE                     KeyOrButton = 0x175
	KeyKEYBOARD                 KeyOrButton = 0x176
	KeyASPECT_RATIO             KeyOrButton = 0x177 // HUTRR37: Aspect
	KeySCREEN                   KeyOrButton = KeyASPECT_RATIO
	KeyPC                       KeyOrButton = 0x178 // Media Select Computer
	KeyTV                       KeyOrButton = 0x179 // Media Select TV
	KeyTV2                      KeyOrButton = 0x17a // Media Select Cable
	KeyVCR                      KeyOrButton = 0x17b // Media Select VCR
	KeyVCR2                     KeyOrButton = 0x17c // VCR Plus
	KeySAT                      KeyOrButton = 0x17d // Media Select Satellite
	KeySAT2                     KeyOrButton = 0x17e
	KeyCD                       KeyOrButton = 0x17f // Media Select CD
	KeyTAPE                     KeyOrButton = 0x180 // Media Select Tape
	KeyRADIO                    KeyOrButton = 0x181
	KeyTUNER                    KeyOrButton = 0x182 // Media Select Tuner
	KeyPLAYER                   KeyOrButton = 0x183
	KeyTEXT                     KeyOrButton = 0x184
	KeyDVD                      KeyOrButton = 0x185 // Media Select DVD
	KeyAUX                      KeyOrButton = 0x186
	KeyMP3                      KeyOrButton = 0x187
	KeyAUDIO                    KeyOrButton = 0x188 // AL Audio Browser
	KeyVIDEO                    KeyOrButton = 0x189 // AL Movie Browser
	KeyDIRECTORY                KeyOrButton = 0x18a
	KeyLIST                     KeyOrButton = 0x18b
	KeyMEMO                     KeyOrButton = 0x18c // Media Select Messages
	KeyCALENDAR                 KeyOrButton = 0x18d
	KeyRED                      KeyOrButton = 0x18e
	KeyGREEN                    KeyOrButton = 0x18f
	KeyYELLOW                   KeyOrButton = 0x190
	KeyBLUE                     KeyOrButton = 0x191
	KeyCHANNELUP                KeyOrButton = 0x192 // Channel Increment
	KeyCHANNELDOWN              KeyOrButton = 0x193 // Channel Decrement
	KeyFIRST                    KeyOrButton = 0x194
	KeyLAST                     KeyOrButton = 0x195 // Recall Last
	KeyAB                       KeyOrButton = 0x196
	KeyNEXT                     KeyOrButton = 0x197
	KeyRESTART                  KeyOrButton = 0x198
	KeySLOW                     KeyOrButton = 0x199
	KeySHUFFLE                  KeyOrButton = 0x19a
	KeyBREAK                    KeyOrButton = 0x19b
	KeyPREVIOUS                 KeyOrButton = 0x19c
	KeyDIGITS                   KeyOrButton = 0x19d
	KeyTEEN                     KeyOrButton = 0x19e
	KeyTWEN                     KeyOrButton = 0x19f
	KeyVIDEOPHONE               KeyOrButton = 0x1a0 // Media Select Video Phone
	KeyGAMES                    KeyOrButton = 0x1a1 // Media Select Games
	KeyZOOMIN                   KeyOrButton = 0x1a2 // AC Zoom In
	KeyZOOMOUT                  KeyOrButton = 0x1a3 // AC Zoom Out
	KeyZOOMRESET                KeyOrButton = 0x1a4 // AC Zoom
	KeyWORDPROCESSOR            KeyOrButton = 0x1a5 // AL Word Processor
	KeyEDITOR                   KeyOrButton = 0x1a6 // AL Text Editor
	KeySPREADSHEET              KeyOrButton = 0x1a7 // AL Spreadsheet
	KeyGRAPHICSEDITOR           KeyOrButton = 0x1a8 // AL Graphics Editor
	KeyPRESENTATION             KeyOrButton = 0x1a9 // AL Presentation App
	KeyDATABASE                 KeyOrButton = 0x1aa // AL Database App
	KeyNEWS                     KeyOrButton = 0x1ab // AL Newsreader
	KeyVOICEMAIL                KeyOrButton = 0x1ac // AL Voicemail
	KeyADDRESSBOOK              KeyOrButton = 0x1ad // AL Contacts/Address Book
	KeyMESSENGER                KeyOrButton = 0x1ae // AL Instant Messaging
	KeyDISPLAYTOGGLE            KeyOrButton = 0x1af // Turn display (LCD) on and off
	KeyBRIGHTNESS_TOGGLE        KeyOrButton = KeyDISPLAYTOGGLE
	KeySPELLCHECK               KeyOrButton = 0x1b0 // AL Spell Check
	KeyLOGOFF                   KeyOrButton = 0x1b1 // AL Logoff
	KeyDOLLAR                   KeyOrButton = 0x1b2
	KeyEURO                     KeyOrButton = 0x1b3
	KeyFRAMEBACK                KeyOrButton = 0x1b4 // Consumer - transport controls
	KeyFRAMEFORWARD             KeyOrButton = 0x1b5
	KeyCONTEXT_MENU             KeyOrButton = 0x1b6 // GenDesc - system context menu
	KeyMEDIA_REPEAT             KeyOrButton = 0x1b7 // Consumer - transport control
	Key10CHANNELSUP             KeyOrButton = 0x1b8 // 10 channels up (10+)
	Key10CHANNELSDOWN           KeyOrButton = 0x1b9 // 10 channels down (10-)
	KeyIMAGES                   KeyOrButton = 0x1ba // AL Image Browser
	KeyNOTIFICATION_CENTER      KeyOrButton = 0x1bc // Show/hide the notification center
	KeyPICKUP_PHONE             KeyOrButton = 0x1bd // Answer incoming call
	KeyHANGUP_PHONE             KeyOrButton = 0x1be // Decline incoming call
	KeyDEL_EOL                  KeyOrButton = 0x1c0
	KeyDEL_EOS                  KeyOrButton = 0x1c1
	KeyINS_LINE                 KeyOrButton = 0x1c2
	KeyDEL_LINE                 KeyOrButton = 0x1c3
	KeyFN                       KeyOrButton = 0x1d0
	KeyFN_ESC                   KeyOrButton = 0x1d1
	KeyFN_F1                    KeyOrButton = 0x1d2
	KeyFN_F2                    KeyOrButton = 0x1d3
	KeyFN_F3                    KeyOrButton = 0x1d4
	KeyFN_F4                    KeyOrButton = 0x1d5
	KeyFN_F5                    KeyOrButton = 0x1d6
	KeyFN_F6                    KeyOrButton = 0x1d7
	KeyFN_F7                    KeyOrButton = 0x1d8
	KeyFN_F8                    KeyOrButton = 0x1d9
	KeyFN_F9                    KeyOrButton = 0x1da
	KeyFN_F10                   KeyOrButton = 0x1db
	KeyFN_F11                   KeyOrButton = 0x1dc
	KeyFN_F12                   KeyOrButton = 0x1dd
	KeyFN_1                     KeyOrButton = 0x1de
	KeyFN_2                     KeyOrButton = 0x1df
	KeyFN_D                     KeyOrButton = 0x1e0
	KeyFN_E                     KeyOrButton = 0x1e1
	KeyFN_F                     KeyOrButton = 0x1e2
	KeyFN_S                     KeyOrButton = 0x1e3
	KeyFN_B                     KeyOrButton = 0x1e4
	KeyFN_RIGHT_SHIFT           KeyOrButton = 0x1e5
	KeyBRL_DOT1                 KeyOrButton = 0x1f1
	KeyBRL_DOT2                 KeyOrButton = 0x1f2
	KeyBRL_DOT3                 KeyOrButton = 0x1f3
	KeyBRL_DOT4                 KeyOrButton = 0x1f4
	KeyBRL_DOT5                 KeyOrButton = 0x1f5
	KeyBRL_DOT6                 KeyOrButton = 0x1f6
	KeyBRL_DOT7                 KeyOrButton = 0x1f7
	KeyBRL_DOT8                 KeyOrButton = 0x1f8
	KeyBRL_DOT9                 KeyOrButton = 0x1f9
	KeyBRL_DOT10                KeyOrButton = 0x1fa
	KeyNUMERIC_0                KeyOrButton = 0x200 // used by phones, remote controls,
	KeyNUMERIC_1                KeyOrButton = 0x201 // and other keypads
	KeyNUMERIC_2                KeyOrButton = 0x202
	KeyNUMERIC_3                KeyOrButton = 0x203
	KeyNUMERIC_4                KeyOrButton = 0x204
	KeyNUMERIC_5                KeyOrButton = 0x205
	KeyNUMERIC_6                KeyOrButton = 0x206
	KeyNUMERIC_7                KeyOrButton = 0x207
	KeyNUMERIC_8                KeyOrButton = 0x208
	KeyNUMERIC_9                KeyOrButton = 0x209
	KeyNUMERIC_STAR             KeyOrButton = 0x20a
	KeyNUMERIC_POUND            KeyOrButton = 0x20b
	KeyNUMERIC_A                KeyOrButton = 0x20c // Phone key A - HUT Telephony 0xb9
	KeyNUMERIC_B                KeyOrButton = 0x20d
	KeyNUMERIC_C                KeyOrButton = 0x20e
	KeyNUMERIC_D                KeyOrButton = 0x20f
	KeyCAMERA_FOCUS             KeyOrButton = 0x210
	KeyWPS_BUTTON               KeyOrButton = 0x211 // WiFi Protected Setup key
	KeyTOUCHPAD_TOGGLE          KeyOrButton = 0x212 // Request switch touchpad on or off
	KeyTOUCHPAD_ON              KeyOrButton = 0x213
	KeyTOUCHPAD_OFF             KeyOrButton = 0x214
	KeyCAMERA_ZOOMIN            KeyOrButton = 0x215
	KeyCAMERA_ZOOMOUT           KeyOrButton = 0x216
	KeyCAMERA_UP                KeyOrButton = 0x217
	KeyCAMERA_DOWN              KeyOrButton = 0x218
	KeyCAMERA_LEFT              KeyOrButton = 0x219
	KeyCAMERA_RIGHT             KeyOrButton = 0x21a
	KeyATTENDANT_ON             KeyOrButton = 0x21b
	KeyATTENDANT_OFF            KeyOrButton = 0x21c
	KeyATTENDANT_TOGGLE         KeyOrButton = 0x21d // Attendant call on or off
	KeyLIGHTS_TOGGLE            KeyOrButton = 0x21e // Reading light on or off
	KeyALS_TOGGLE               KeyOrButton = 0x230 // Ambient light sensor
	KeyROTATE_LOCK_TOGGLE       KeyOrButton = 0x231 // Display rotation lock
	KeyBUTTONCONFIG             KeyOrButton = 0x240 // AL Button Configuration
	KeyTASKMANAGER              KeyOrButton = 0x241 // AL Task/Project Manager
	KeyJOURNAL                  KeyOrButton = 0x242 // AL Log/Journal/Timecard
	KeyCONTROLPANEL             KeyOrButton = 0x243 // AL Control Panel
	KeyAPPSELECT                KeyOrButton = 0x244 // AL Select Task/Application
	KeySCREENSAVER              KeyOrButton = 0x245 // AL Screen Saver
	KeyVOICECOMMAND             KeyOrButton = 0x246 // Listening Voice Command
	KeyASSISTANT                KeyOrButton = 0x247 // AL Context-aware desktop assistant
	KeyKBD_LAYOUT_NEXT          KeyOrButton = 0x248 // AC Next Keyboard Layout Select
	KeyBRIGHTNESS_MIN           KeyOrButton = 0x250 // Set Brightness to Minimum
	KeyBRIGHTNESS_MAX           KeyOrButton = 0x251 // Set Brightness to Maximum
	KeyKBDINPUTASSIST_PREV      KeyOrButton = 0x260
	KeyKBDINPUTASSIST_NEXT      KeyOrButton = 0x261
	KeyKBDINPUTASSIST_PREVGROUP KeyOrButton = 0x262
	KeyKBDINPUTASSIST_NEXTGROUP KeyOrButton = 0x263
	KeyKBDINPUTASSIST_ACCEPT    KeyOrButton = 0x264
	KeyKBDINPUTASSIST_CANCEL    KeyOrButton = 0x265
	KeyRIGHT_UP                 KeyOrButton = 0x266
	KeyRIGHT_DOWN               KeyOrButton = 0x267
	KeyLEFT_UP                  KeyOrButton = 0x268
	KeyLEFT_DOWN                KeyOrButton = 0x269
	KeyROOT_MENU                KeyOrButton = 0x26a // Show Device's Root Menu
	KeyMEDIA_TOP_MENU           KeyOrButton = 0x26b
	KeyNUMERIC_11               KeyOrButton = 0x26c
	KeyNUMERIC_12               KeyOrButton = 0x26d
	KeyAUDIO_DESC               KeyOrButton = 0x26e
	Key3D_MODE                  KeyOrButton = 0x26f
	KeyNEXT_FAVORITE            KeyOrButton = 0x270
	KeySTOP_RECORD              KeyOrButton = 0x271
	KeyPAUSE_RECORD             KeyOrButton = 0x272
	KeyVOD                      KeyOrButton = 0x273 // Video on Demand
	KeyUNMUTE                   KeyOrButton = 0x274
	KeyFASTREVERSE              KeyOrButton = 0x275
	KeySLOWREVERSE              KeyOrButton = 0x276
	KeyDATA                     KeyOrButton = 0x277
	KeyONSCREEN_KEYBOARD        KeyOrButton = 0x278
	KeyPRIVACY_SCREEN_TOGGLE    KeyOrButton = 0x279
	KeySELECTIVE_SCREENSHOT     KeyOrButton = 0x27a
	KeyMACRO1                   KeyOrButton = 0x290
	KeyMACRO2                   KeyOrButton = 0x291
	KeyMACRO3                   KeyOrButton = 0x292
	KeyMACRO4                   KeyOrButton = 0x293
	KeyMACRO5                   KeyOrButton = 0x294
	KeyMACRO6                   KeyOrButton = 0x295
	KeyMACRO7                   KeyOrButton = 0x296
	KeyMACRO8                   KeyOrButton = 0x297
	KeyMACRO9                   KeyOrButton = 0x298
	KeyMACRO10                  KeyOrButton = 0x299
	KeyMACRO11                  KeyOrButton = 0x29a
	KeyMACRO12                  KeyOrButton = 0x29b
	KeyMACRO13                  KeyOrButton = 0x29c
	KeyMACRO14                  KeyOrButton = 0x29d
	KeyMACRO15                  KeyOrButton = 0x29e
	KeyMACRO16                  KeyOrButton = 0x29f
	KeyMACRO17                  KeyOrButton = 0x2a0
	KeyMACRO18                  KeyOrButton = 0x2a1
	KeyMACRO19                  KeyOrButton = 0x2a2
	KeyMACRO20                  KeyOrButton = 0x2a3
	KeyMACRO21                  KeyOrButton = 0x2a4
	KeyMACRO22                  KeyOrButton = 0x2a5
	KeyMACRO23                  KeyOrButton = 0x2a6
	KeyMACRO24                  KeyOrButton = 0x2a7
	KeyMACRO25                  KeyOrButton = 0x2a8
	KeyMACRO26                  KeyOrButton = 0x2a9
	KeyMACRO27                  KeyOrButton = 0x2aa
	KeyMACRO28                  KeyOrButton = 0x2ab
	KeyMACRO29                  KeyOrButton = 0x2ac
	KeyMACRO30                  KeyOrButton = 0x2ad
	KeyMACRO_RECORD_START       KeyOrButton = 0x2b0
	KeyMACRO_RECORD_STOP        KeyOrButton = 0x2b1
	KeyMACRO_PRESET_CYCLE       KeyOrButton = 0x2b2
	KeyMACRO_PRESET1            KeyOrButton = 0x2b3
	KeyMACRO_PRESET2            KeyOrButton = 0x2b4
	KeyMACRO_PRESET3            KeyOrButton = 0x2b5
	KeyKBD_LCD_MENU1            KeyOrButton = 0x2b8
	KeyKBD_LCD_MENU2            KeyOrButton = 0x2b9
	KeyKBD_LCD_MENU3            KeyOrButton = 0x2ba
	KeyKBD_LCD_MENU4            KeyOrButton = 0x2bb
	KeyKBD_LCD_MENU5            KeyOrButton = 0x2bc
	KeyMIN_INTERESTING          KeyOrButton = KeyMUTE
)

// the below are in evdev definitions with buttons prefix.
// we merge them here under the the same type *KeyOrButton* because:
// - we can: they have distinct number namespaces anyway
// - I don't see a clear reason for separating "key" and "button" namespaces..
//   they're all switches with some kind of label?

// https://github.com/torvalds/linux/blob/5c8fe583cc/include/uapi/linux/input-event-codes.h#L342

const (
	BtnMISC            KeyOrButton = Btn0 // alias
	Btn0               KeyOrButton = 0x100
	Btn1               KeyOrButton = 0x101
	Btn2               KeyOrButton = 0x102
	Btn3               KeyOrButton = 0x103
	Btn4               KeyOrButton = 0x104
	Btn5               KeyOrButton = 0x105
	Btn6               KeyOrButton = 0x106
	Btn7               KeyOrButton = 0x107
	Btn8               KeyOrButton = 0x108
	Btn9               KeyOrButton = 0x109
	BtnMOUSE           KeyOrButton = BtnLEFT // alias
	BtnLEFT            KeyOrButton = 0x110
	BtnRIGHT           KeyOrButton = 0x111
	BtnMIDDLE          KeyOrButton = 0x112
	BtnSIDE            KeyOrButton = 0x113
	BtnEXTRA           KeyOrButton = 0x114
	BtnFORWARD         KeyOrButton = 0x115
	BtnBACK            KeyOrButton = 0x116
	BtnTASK            KeyOrButton = 0x117
	BtnJOYSTICK        KeyOrButton = BtnTRIGGER // alias
	BtnTRIGGER         KeyOrButton = 0x120
	BtnTHUMB           KeyOrButton = 0x121
	BtnTHUMB2          KeyOrButton = 0x122
	BtnTOP             KeyOrButton = 0x123
	BtnTOP2            KeyOrButton = 0x124
	BtnPINKIE          KeyOrButton = 0x125
	BtnBASE            KeyOrButton = 0x126
	BtnBASE2           KeyOrButton = 0x127
	BtnBASE3           KeyOrButton = 0x128
	BtnBASE4           KeyOrButton = 0x129
	BtnBASE5           KeyOrButton = 0x12a
	BtnBASE6           KeyOrButton = 0x12b
	BtnDEAD            KeyOrButton = 0x12f
	BtnGAMEPAD         KeyOrButton = BtnSOUTH // alias
	BtnSOUTH           KeyOrButton = 0x130
	BtnA               KeyOrButton = BtnSOUTH
	BtnEAST            KeyOrButton = 0x131
	BtnB               KeyOrButton = BtnEAST
	BtnC               KeyOrButton = 0x132
	BtnNORTH           KeyOrButton = 0x133
	BtnX               KeyOrButton = BtnNORTH
	BtnWEST            KeyOrButton = 0x134
	BtnY               KeyOrButton = BtnWEST
	BtnZ               KeyOrButton = 0x135
	BtnTL              KeyOrButton = 0x136
	BtnTR              KeyOrButton = 0x137
	BtnTL2             KeyOrButton = 0x138
	BtnTR2             KeyOrButton = 0x139
	BtnSELECT          KeyOrButton = 0x13a
	BtnSTART           KeyOrButton = 0x13b
	BtnMODE            KeyOrButton = 0x13c
	BtnTHUMBL          KeyOrButton = 0x13d
	BtnTHUMBR          KeyOrButton = 0x13e
	BtnDIGI            KeyOrButton = BtnTOOL_PEN // alias
	BtnTOOL_PEN        KeyOrButton = 0x140
	BtnTOOL_RUBBER     KeyOrButton = 0x141
	BtnTOOL_BRUSH      KeyOrButton = 0x142
	BtnTOOL_PENCIL     KeyOrButton = 0x143
	BtnTOOL_AIRBRUSH   KeyOrButton = 0x144
	BtnTOOL_FINGER     KeyOrButton = 0x145
	BtnTOOL_MOUSE      KeyOrButton = 0x146
	BtnTOOL_LENS       KeyOrButton = 0x147
	BtnTOOL_QUINTTAP   KeyOrButton = 0x148 // Five fingers on trackpad
	BtnSTYLUS3         KeyOrButton = 0x149
	BtnTOUCH           KeyOrButton = 0x14a
	BtnSTYLUS          KeyOrButton = 0x14b
	BtnSTYLUS2         KeyOrButton = 0x14c
	BtnTOOL_DOUBLETAP  KeyOrButton = 0x14d
	BtnTOOL_TRIPLETAP  KeyOrButton = 0x14e
	BtnTOOL_QUADTAP    KeyOrButton = 0x14f        // Four fingers on trackpad
	BtnWHEEL           KeyOrButton = BtnGEAR_DOWN // alias
	BtnGEAR_DOWN       KeyOrButton = 0x150
	BtnGEAR_UP         KeyOrButton = 0x151
	BtnDPAD_UP         KeyOrButton = 0x220
	BtnDPAD_DOWN       KeyOrButton = 0x221
	BtnDPAD_LEFT       KeyOrButton = 0x222
	BtnDPAD_RIGHT      KeyOrButton = 0x223
	BtnTRIGGER_HAPPY   KeyOrButton = BtnTRIGGER_HAPPY1 // alias
	BtnTRIGGER_HAPPY1  KeyOrButton = 0x2c0
	BtnTRIGGER_HAPPY2  KeyOrButton = 0x2c1
	BtnTRIGGER_HAPPY3  KeyOrButton = 0x2c2
	BtnTRIGGER_HAPPY4  KeyOrButton = 0x2c3
	BtnTRIGGER_HAPPY5  KeyOrButton = 0x2c4
	BtnTRIGGER_HAPPY6  KeyOrButton = 0x2c5
	BtnTRIGGER_HAPPY7  KeyOrButton = 0x2c6
	BtnTRIGGER_HAPPY8  KeyOrButton = 0x2c7
	BtnTRIGGER_HAPPY9  KeyOrButton = 0x2c8
	BtnTRIGGER_HAPPY10 KeyOrButton = 0x2c9
	BtnTRIGGER_HAPPY11 KeyOrButton = 0x2ca
	BtnTRIGGER_HAPPY12 KeyOrButton = 0x2cb
	BtnTRIGGER_HAPPY13 KeyOrButton = 0x2cc
	BtnTRIGGER_HAPPY14 KeyOrButton = 0x2cd
	BtnTRIGGER_HAPPY15 KeyOrButton = 0x2ce
	BtnTRIGGER_HAPPY16 KeyOrButton = 0x2cf
	BtnTRIGGER_HAPPY17 KeyOrButton = 0x2d0
	BtnTRIGGER_HAPPY18 KeyOrButton = 0x2d1
	BtnTRIGGER_HAPPY19 KeyOrButton = 0x2d2
	BtnTRIGGER_HAPPY20 KeyOrButton = 0x2d3
	BtnTRIGGER_HAPPY21 KeyOrButton = 0x2d4
	BtnTRIGGER_HAPPY22 KeyOrButton = 0x2d5
	BtnTRIGGER_HAPPY23 KeyOrButton = 0x2d6
	BtnTRIGGER_HAPPY24 KeyOrButton = 0x2d7
	BtnTRIGGER_HAPPY25 KeyOrButton = 0x2d8
	BtnTRIGGER_HAPPY26 KeyOrButton = 0x2d9
	BtnTRIGGER_HAPPY27 KeyOrButton = 0x2da
	BtnTRIGGER_HAPPY28 KeyOrButton = 0x2db
	BtnTRIGGER_HAPPY29 KeyOrButton = 0x2dc
	BtnTRIGGER_HAPPY30 KeyOrButton = 0x2dd
	BtnTRIGGER_HAPPY31 KeyOrButton = 0x2de
	BtnTRIGGER_HAPPY32 KeyOrButton = 0x2df
	BtnTRIGGER_HAPPY33 KeyOrButton = 0x2e0
	BtnTRIGGER_HAPPY34 KeyOrButton = 0x2e1
	BtnTRIGGER_HAPPY35 KeyOrButton = 0x2e2
	BtnTRIGGER_HAPPY36 KeyOrButton = 0x2e3
	BtnTRIGGER_HAPPY37 KeyOrButton = 0x2e4
	BtnTRIGGER_HAPPY38 KeyOrButton = 0x2e5
	BtnTRIGGER_HAPPY39 KeyOrButton = 0x2e6
	BtnTRIGGER_HAPPY40 KeyOrButton = 0x2e7
)

var keyToString = map[KeyOrButton]string{
	KeyRESERVED:                 "RESERVED",
	KeyESC:                      "ESC",
	Key1:                        "1",
	Key2:                        "2",
	Key3:                        "3",
	Key4:                        "4",
	Key5:                        "5",
	Key6:                        "6",
	Key7:                        "7",
	Key8:                        "8",
	Key9:                        "9",
	Key0:                        "0",
	KeyMINUS:                    "MINUS",
	KeyEQUAL:                    "EQUAL",
	KeyBACKSPACE:                "BACKSPACE",
	KeyTAB:                      "TAB",
	KeyQ:                        "Q",
	KeyW:                        "W",
	KeyE:                        "E",
	KeyR:                        "R",
	KeyT:                        "T",
	KeyY:                        "Y",
	KeyU:                        "U",
	KeyI:                        "I",
	KeyO:                        "O",
	KeyP:                        "P",
	KeyLEFTBRACE:                "LEFTBRACE",
	KeyRIGHTBRACE:               "RIGHTBRACE",
	KeyENTER:                    "ENTER",
	KeyLEFTCTRL:                 "LEFTCTRL",
	KeyA:                        "A",
	KeyS:                        "S",
	KeyD:                        "D",
	KeyF:                        "F",
	KeyG:                        "G",
	KeyH:                        "H",
	KeyJ:                        "J",
	KeyK:                        "K",
	KeyL:                        "L",
	KeySEMICOLON:                "SEMICOLON",
	KeyAPOSTROPHE:               "APOSTROPHE",
	KeyGRAVE:                    "GRAVE",
	KeyLEFTSHIFT:                "LEFTSHIFT",
	KeyBACKSLASH:                "BACKSLASH",
	KeyZ:                        "Z",
	KeyX:                        "X",
	KeyC:                        "C",
	KeyV:                        "V",
	KeyB:                        "B",
	KeyN:                        "N",
	KeyM:                        "M",
	KeyCOMMA:                    "COMMA",
	KeyDOT:                      "DOT",
	KeySLASH:                    "SLASH",
	KeyRIGHTSHIFT:               "RIGHTSHIFT",
	KeyKPASTERISK:               "KPASTERISK",
	KeyLEFTALT:                  "LEFTALT",
	KeySPACE:                    "SPACE",
	KeyCAPSLOCK:                 "CAPSLOCK",
	KeyF1:                       "F1",
	KeyF2:                       "F2",
	KeyF3:                       "F3",
	KeyF4:                       "F4",
	KeyF5:                       "F5",
	KeyF6:                       "F6",
	KeyF7:                       "F7",
	KeyF8:                       "F8",
	KeyF9:                       "F9",
	KeyF10:                      "F10",
	KeyNUMLOCK:                  "NUMLOCK",
	KeySCROLLLOCK:               "SCROLLLOCK",
	KeyKP7:                      "KP7",
	KeyKP8:                      "KP8",
	KeyKP9:                      "KP9",
	KeyKPMINUS:                  "KPMINUS",
	KeyKP4:                      "KP4",
	KeyKP5:                      "KP5",
	KeyKP6:                      "KP6",
	KeyKPPLUS:                   "KPPLUS",
	KeyKP1:                      "KP1",
	KeyKP2:                      "KP2",
	KeyKP3:                      "KP3",
	KeyKP0:                      "KP0",
	KeyKPDOT:                    "KPDOT",
	KeyZENKAKUHANKAKU:           "ZENKAKUHANKAKU",
	Key102ND:                    "102ND",
	KeyF11:                      "F11",
	KeyF12:                      "F12",
	KeyRO:                       "RO",
	KeyKATAKANA:                 "KATAKANA",
	KeyHIRAGANA:                 "HIRAGANA",
	KeyHENKAN:                   "HENKAN",
	KeyKATAKANAHIRAGANA:         "KATAKANAHIRAGANA",
	KeyMUHENKAN:                 "MUHENKAN",
	KeyKPJPCOMMA:                "KPJPCOMMA",
	KeyKPENTER:                  "KPENTER",
	KeyRIGHTCTRL:                "RIGHTCTRL",
	KeyKPSLASH:                  "KPSLASH",
	KeySYSRQ:                    "SYSRQ",
	KeyRIGHTALT:                 "RIGHTALT",
	KeyLINEFEED:                 "LINEFEED",
	KeyHOME:                     "HOME",
	KeyUP:                       "UP",
	KeyPAGEUP:                   "PAGEUP",
	KeyLEFT:                     "LEFT",
	KeyRIGHT:                    "RIGHT",
	KeyEND:                      "END",
	KeyDOWN:                     "DOWN",
	KeyPAGEDOWN:                 "PAGEDOWN",
	KeyINSERT:                   "INSERT",
	KeyDELETE:                   "DELETE",
	KeyMACRO:                    "MACRO",
	KeyMUTE:                     "MUTE",
	KeyVOLUMEDOWN:               "VOLUMEDOWN",
	KeyVOLUMEUP:                 "VOLUMEUP",
	KeyPOWER:                    "POWER",
	KeyKPEQUAL:                  "KPEQUAL",
	KeyKPPLUSMINUS:              "KPPLUSMINUS",
	KeyPAUSE:                    "PAUSE",
	KeySCALE:                    "SCALE",
	KeyKPCOMMA:                  "KPCOMMA",
	KeyHANGEUL:                  "HANGEUL",
	KeyHANJA:                    "HANJA",
	KeyYEN:                      "YEN",
	KeyLEFTMETA:                 "LEFTMETA",
	KeyRIGHTMETA:                "RIGHTMETA",
	KeyCOMPOSE:                  "COMPOSE",
	KeySTOP:                     "STOP",
	KeyAGAIN:                    "AGAIN",
	KeyPROPS:                    "PROPS",
	KeyUNDO:                     "UNDO",
	KeyFRONT:                    "FRONT",
	KeyCOPY:                     "COPY",
	KeyOPEN:                     "OPEN",
	KeyPASTE:                    "PASTE",
	KeyFIND:                     "FIND",
	KeyCUT:                      "CUT",
	KeyHELP:                     "HELP",
	KeyMENU:                     "MENU",
	KeyCALC:                     "CALC",
	KeySETUP:                    "SETUP",
	KeySLEEP:                    "SLEEP",
	KeyWAKEUP:                   "WAKEUP",
	KeyFILE:                     "FILE",
	KeySENDFILE:                 "SENDFILE",
	KeyDELETEFILE:               "DELETEFILE",
	KeyXFER:                     "XFER",
	KeyPROG1:                    "PROG1",
	KeyPROG2:                    "PROG2",
	KeyWWW:                      "WWW",
	KeyMSDOS:                    "MSDOS",
	KeyCOFFEE:                   "COFFEE",
	KeyROTATE_DISPLAY:           "ROTATE_DISPLAY",
	KeyCYCLEWINDOWS:             "CYCLEWINDOWS",
	KeyMAIL:                     "MAIL",
	KeyBOOKMARKS:                "BOOKMARKS",
	KeyCOMPUTER:                 "COMPUTER",
	KeyBACK:                     "BACK",
	KeyFORWARD:                  "FORWARD",
	KeyCLOSECD:                  "CLOSECD",
	KeyEJECTCD:                  "EJECTCD",
	KeyEJECTCLOSECD:             "EJECTCLOSECD",
	KeyNEXTSONG:                 "NEXTSONG",
	KeyPLAYPAUSE:                "PLAYPAUSE",
	KeyPREVIOUSSONG:             "PREVIOUSSONG",
	KeySTOPCD:                   "STOPCD",
	KeyRECORD:                   "RECORD",
	KeyREWIND:                   "REWIND",
	KeyPHONE:                    "PHONE",
	KeyISO:                      "ISO",
	KeyCONFIG:                   "CONFIG",
	KeyHOMEPAGE:                 "HOMEPAGE",
	KeyREFRESH:                  "REFRESH",
	KeyEXIT:                     "EXIT",
	KeyMOVE:                     "MOVE",
	KeyEDIT:                     "EDIT",
	KeySCROLLUP:                 "SCROLLUP",
	KeySCROLLDOWN:               "SCROLLDOWN",
	KeyKPLEFTPAREN:              "KPLEFTPAREN",
	KeyKPRIGHTPAREN:             "KPRIGHTPAREN",
	KeyNEW:                      "NEW",
	KeyREDO:                     "REDO",
	KeyF13:                      "F13",
	KeyF14:                      "F14",
	KeyF15:                      "F15",
	KeyF16:                      "F16",
	KeyF17:                      "F17",
	KeyF18:                      "F18",
	KeyF19:                      "F19",
	KeyF20:                      "F20",
	KeyF21:                      "F21",
	KeyF22:                      "F22",
	KeyF23:                      "F23",
	KeyF24:                      "F24",
	KeyPLAYCD:                   "PLAYCD",
	KeyPAUSECD:                  "PAUSECD",
	KeyPROG3:                    "PROG3",
	KeyPROG4:                    "PROG4",
	KeyDASHBOARD:                "DASHBOARD",
	KeySUSPEND:                  "SUSPEND",
	KeyCLOSE:                    "CLOSE",
	KeyPLAY:                     "PLAY",
	KeyFASTFORWARD:              "FASTFORWARD",
	KeyBASSBOOST:                "BASSBOOST",
	KeyPRINT:                    "PRINT",
	KeyHP:                       "HP",
	KeyCAMERA:                   "CAMERA",
	KeySOUND:                    "SOUND",
	KeyQUESTION:                 "QUESTION",
	KeyEMAIL:                    "EMAIL",
	KeyCHAT:                     "CHAT",
	KeySEARCH:                   "SEARCH",
	KeyCONNECT:                  "CONNECT",
	KeyFINANCE:                  "FINANCE",
	KeySPORT:                    "SPORT",
	KeySHOP:                     "SHOP",
	KeyALTERASE:                 "ALTERASE",
	KeyCANCEL:                   "CANCEL",
	KeyBRIGHTNESSDOWN:           "BRIGHTNESSDOWN",
	KeyBRIGHTNESSUP:             "BRIGHTNESSUP",
	KeyMEDIA:                    "MEDIA",
	KeySWITCHVIDEOMODE:          "SWITCHVIDEOMODE",
	KeyKBDILLUMTOGGLE:           "KBDILLUMTOGGLE",
	KeyKBDILLUMDOWN:             "KBDILLUMDOWN",
	KeyKBDILLUMUP:               "KBDILLUMUP",
	KeySEND:                     "SEND",
	KeyREPLY:                    "REPLY",
	KeyFORWARDMAIL:              "FORWARDMAIL",
	KeySAVE:                     "SAVE",
	KeyDOCUMENTS:                "DOCUMENTS",
	KeyBATTERY:                  "BATTERY",
	KeyBLUETOOTH:                "BLUETOOTH",
	KeyWLAN:                     "WLAN",
	KeyUWB:                      "UWB",
	KeyUNKNOWN:                  "UNKNOWN",
	KeyVIDEO_NEXT:               "VIDEO_NEXT",
	KeyVIDEO_PREV:               "VIDEO_PREV",
	KeyBRIGHTNESS_CYCLE:         "BRIGHTNESS_CYCLE",
	KeyBRIGHTNESS_AUTO:          "BRIGHTNESS_AUTO",
	KeyDISPLAY_OFF:              "DISPLAY_OFF",
	KeyWWAN:                     "WWAN",
	KeyRFKILL:                   "RFKILL",
	KeyMICMUTE:                  "MICMUTE",
	KeyOK:                       "OK",
	KeySELECT:                   "SELECT",
	KeyGOTO:                     "GOTO",
	KeyCLEAR:                    "CLEAR",
	KeyPOWER2:                   "POWER2",
	KeyOPTION:                   "OPTION",
	KeyINFO:                     "INFO",
	KeyTIME:                     "TIME",
	KeyVENDOR:                   "VENDOR",
	KeyARCHIVE:                  "ARCHIVE",
	KeyPROGRAM:                  "PROGRAM",
	KeyCHANNEL:                  "CHANNEL",
	KeyFAVORITES:                "FAVORITES",
	KeyEPG:                      "EPG",
	KeyPVR:                      "PVR",
	KeyMHP:                      "MHP",
	KeyLANGUAGE:                 "LANGUAGE",
	KeyTITLE:                    "TITLE",
	KeySUBTITLE:                 "SUBTITLE",
	KeyANGLE:                    "ANGLE",
	KeyFULL_SCREEN:              "FULL_SCREEN",
	KeyMODE:                     "MODE",
	KeyKEYBOARD:                 "KEYBOARD",
	KeyASPECT_RATIO:             "ASPECT_RATIO",
	KeyPC:                       "PC",
	KeyTV:                       "TV",
	KeyTV2:                      "TV2",
	KeyVCR:                      "VCR",
	KeyVCR2:                     "VCR2",
	KeySAT:                      "SAT",
	KeySAT2:                     "SAT2",
	KeyCD:                       "CD",
	KeyTAPE:                     "TAPE",
	KeyRADIO:                    "RADIO",
	KeyTUNER:                    "TUNER",
	KeyPLAYER:                   "PLAYER",
	KeyTEXT:                     "TEXT",
	KeyDVD:                      "DVD",
	KeyAUX:                      "AUX",
	KeyMP3:                      "MP3",
	KeyAUDIO:                    "AUDIO",
	KeyVIDEO:                    "VIDEO",
	KeyDIRECTORY:                "DIRECTORY",
	KeyLIST:                     "LIST",
	KeyMEMO:                     "MEMO",
	KeyCALENDAR:                 "CALENDAR",
	KeyRED:                      "RED",
	KeyGREEN:                    "GREEN",
	KeyYELLOW:                   "YELLOW",
	KeyBLUE:                     "BLUE",
	KeyCHANNELUP:                "CHANNELUP",
	KeyCHANNELDOWN:              "CHANNELDOWN",
	KeyFIRST:                    "FIRST",
	KeyLAST:                     "LAST",
	KeyAB:                       "AB",
	KeyNEXT:                     "NEXT",
	KeyRESTART:                  "RESTART",
	KeySLOW:                     "SLOW",
	KeySHUFFLE:                  "SHUFFLE",
	KeyBREAK:                    "BREAK",
	KeyPREVIOUS:                 "PREVIOUS",
	KeyDIGITS:                   "DIGITS",
	KeyTEEN:                     "TEEN",
	KeyTWEN:                     "TWEN",
	KeyVIDEOPHONE:               "VIDEOPHONE",
	KeyGAMES:                    "GAMES",
	KeyZOOMIN:                   "ZOOMIN",
	KeyZOOMOUT:                  "ZOOMOUT",
	KeyZOOMRESET:                "ZOOMRESET",
	KeyWORDPROCESSOR:            "WORDPROCESSOR",
	KeyEDITOR:                   "EDITOR",
	KeySPREADSHEET:              "SPREADSHEET",
	KeyGRAPHICSEDITOR:           "GRAPHICSEDITOR",
	KeyPRESENTATION:             "PRESENTATION",
	KeyDATABASE:                 "DATABASE",
	KeyNEWS:                     "NEWS",
	KeyVOICEMAIL:                "VOICEMAIL",
	KeyADDRESSBOOK:              "ADDRESSBOOK",
	KeyMESSENGER:                "MESSENGER",
	KeyDISPLAYTOGGLE:            "DISPLAYTOGGLE",
	KeySPELLCHECK:               "SPELLCHECK",
	KeyLOGOFF:                   "LOGOFF",
	KeyDOLLAR:                   "DOLLAR",
	KeyEURO:                     "EURO",
	KeyFRAMEBACK:                "FRAMEBACK",
	KeyFRAMEFORWARD:             "FRAMEFORWARD",
	KeyCONTEXT_MENU:             "CONTEXT_MENU",
	KeyMEDIA_REPEAT:             "MEDIA_REPEAT",
	Key10CHANNELSUP:             "10CHANNELSUP",
	Key10CHANNELSDOWN:           "10CHANNELSDOWN",
	KeyIMAGES:                   "IMAGES",
	KeyNOTIFICATION_CENTER:      "NOTIFICATION_CENTER",
	KeyPICKUP_PHONE:             "PICKUP_PHONE",
	KeyHANGUP_PHONE:             "HANGUP_PHONE",
	KeyDEL_EOL:                  "DEL_EOL",
	KeyDEL_EOS:                  "DEL_EOS",
	KeyINS_LINE:                 "INS_LINE",
	KeyDEL_LINE:                 "DEL_LINE",
	KeyFN:                       "FN",
	KeyFN_ESC:                   "FN_ESC",
	KeyFN_F1:                    "FN_F1",
	KeyFN_F2:                    "FN_F2",
	KeyFN_F3:                    "FN_F3",
	KeyFN_F4:                    "FN_F4",
	KeyFN_F5:                    "FN_F5",
	KeyFN_F6:                    "FN_F6",
	KeyFN_F7:                    "FN_F7",
	KeyFN_F8:                    "FN_F8",
	KeyFN_F9:                    "FN_F9",
	KeyFN_F10:                   "FN_F10",
	KeyFN_F11:                   "FN_F11",
	KeyFN_F12:                   "FN_F12",
	KeyFN_1:                     "FN_1",
	KeyFN_2:                     "FN_2",
	KeyFN_D:                     "FN_D",
	KeyFN_E:                     "FN_E",
	KeyFN_F:                     "FN_F",
	KeyFN_S:                     "FN_S",
	KeyFN_B:                     "FN_B",
	KeyFN_RIGHT_SHIFT:           "FN_RIGHT_SHIFT",
	KeyBRL_DOT1:                 "BRL_DOT1",
	KeyBRL_DOT2:                 "BRL_DOT2",
	KeyBRL_DOT3:                 "BRL_DOT3",
	KeyBRL_DOT4:                 "BRL_DOT4",
	KeyBRL_DOT5:                 "BRL_DOT5",
	KeyBRL_DOT6:                 "BRL_DOT6",
	KeyBRL_DOT7:                 "BRL_DOT7",
	KeyBRL_DOT8:                 "BRL_DOT8",
	KeyBRL_DOT9:                 "BRL_DOT9",
	KeyBRL_DOT10:                "BRL_DOT10",
	KeyNUMERIC_0:                "NUMERIC_0",
	KeyNUMERIC_1:                "NUMERIC_1",
	KeyNUMERIC_2:                "NUMERIC_2",
	KeyNUMERIC_3:                "NUMERIC_3",
	KeyNUMERIC_4:                "NUMERIC_4",
	KeyNUMERIC_5:                "NUMERIC_5",
	KeyNUMERIC_6:                "NUMERIC_6",
	KeyNUMERIC_7:                "NUMERIC_7",
	KeyNUMERIC_8:                "NUMERIC_8",
	KeyNUMERIC_9:                "NUMERIC_9",
	KeyNUMERIC_STAR:             "NUMERIC_STAR",
	KeyNUMERIC_POUND:            "NUMERIC_POUND",
	KeyNUMERIC_A:                "NUMERIC_A",
	KeyNUMERIC_B:                "NUMERIC_B",
	KeyNUMERIC_C:                "NUMERIC_C",
	KeyNUMERIC_D:                "NUMERIC_D",
	KeyCAMERA_FOCUS:             "CAMERA_FOCUS",
	KeyWPS_BUTTON:               "WPS_BUTTON",
	KeyTOUCHPAD_TOGGLE:          "TOUCHPAD_TOGGLE",
	KeyTOUCHPAD_ON:              "TOUCHPAD_ON",
	KeyTOUCHPAD_OFF:             "TOUCHPAD_OFF",
	KeyCAMERA_ZOOMIN:            "CAMERA_ZOOMIN",
	KeyCAMERA_ZOOMOUT:           "CAMERA_ZOOMOUT",
	KeyCAMERA_UP:                "CAMERA_UP",
	KeyCAMERA_DOWN:              "CAMERA_DOWN",
	KeyCAMERA_LEFT:              "CAMERA_LEFT",
	KeyCAMERA_RIGHT:             "CAMERA_RIGHT",
	KeyATTENDANT_ON:             "ATTENDANT_ON",
	KeyATTENDANT_OFF:            "ATTENDANT_OFF",
	KeyATTENDANT_TOGGLE:         "ATTENDANT_TOGGLE",
	KeyLIGHTS_TOGGLE:            "LIGHTS_TOGGLE",
	KeyALS_TOGGLE:               "ALS_TOGGLE",
	KeyROTATE_LOCK_TOGGLE:       "ROTATE_LOCK_TOGGLE",
	KeyBUTTONCONFIG:             "BUTTONCONFIG",
	KeyTASKMANAGER:              "TASKMANAGER",
	KeyJOURNAL:                  "JOURNAL",
	KeyCONTROLPANEL:             "CONTROLPANEL",
	KeyAPPSELECT:                "APPSELECT",
	KeySCREENSAVER:              "SCREENSAVER",
	KeyVOICECOMMAND:             "VOICECOMMAND",
	KeyASSISTANT:                "ASSISTANT",
	KeyKBD_LAYOUT_NEXT:          "KBD_LAYOUT_NEXT",
	KeyBRIGHTNESS_MIN:           "BRIGHTNESS_MIN",
	KeyBRIGHTNESS_MAX:           "BRIGHTNESS_MAX",
	KeyKBDINPUTASSIST_PREV:      "KBDINPUTASSIST_PREV",
	KeyKBDINPUTASSIST_NEXT:      "KBDINPUTASSIST_NEXT",
	KeyKBDINPUTASSIST_PREVGROUP: "KBDINPUTASSIST_PREVGROUP",
	KeyKBDINPUTASSIST_NEXTGROUP: "KBDINPUTASSIST_NEXTGROUP",
	KeyKBDINPUTASSIST_ACCEPT:    "KBDINPUTASSIST_ACCEPT",
	KeyKBDINPUTASSIST_CANCEL:    "KBDINPUTASSIST_CANCEL",
	KeyRIGHT_UP:                 "RIGHT_UP",
	KeyRIGHT_DOWN:               "RIGHT_DOWN",
	KeyLEFT_UP:                  "LEFT_UP",
	KeyLEFT_DOWN:                "LEFT_DOWN",
	KeyROOT_MENU:                "ROOT_MENU",
	KeyMEDIA_TOP_MENU:           "MEDIA_TOP_MENU",
	KeyNUMERIC_11:               "NUMERIC_11",
	KeyNUMERIC_12:               "NUMERIC_12",
	KeyAUDIO_DESC:               "AUDIO_DESC",
	Key3D_MODE:                  "3D_MODE",
	KeyNEXT_FAVORITE:            "NEXT_FAVORITE",
	KeySTOP_RECORD:              "STOP_RECORD",
	KeyPAUSE_RECORD:             "PAUSE_RECORD",
	KeyVOD:                      "VOD",
	KeyUNMUTE:                   "UNMUTE",
	KeyFASTREVERSE:              "FASTREVERSE",
	KeySLOWREVERSE:              "SLOWREVERSE",
	KeyDATA:                     "DATA",
	KeyONSCREEN_KEYBOARD:        "ONSCREEN_KEYBOARD",
	KeyPRIVACY_SCREEN_TOGGLE:    "PRIVACY_SCREEN_TOGGLE",
	KeySELECTIVE_SCREENSHOT:     "SELECTIVE_SCREENSHOT",
	KeyMACRO1:                   "MACRO1",
	KeyMACRO2:                   "MACRO2",
	KeyMACRO3:                   "MACRO3",
	KeyMACRO4:                   "MACRO4",
	KeyMACRO5:                   "MACRO5",
	KeyMACRO6:                   "MACRO6",
	KeyMACRO7:                   "MACRO7",
	KeyMACRO8:                   "MACRO8",
	KeyMACRO9:                   "MACRO9",
	KeyMACRO10:                  "MACRO10",
	KeyMACRO11:                  "MACRO11",
	KeyMACRO12:                  "MACRO12",
	KeyMACRO13:                  "MACRO13",
	KeyMACRO14:                  "MACRO14",
	KeyMACRO15:                  "MACRO15",
	KeyMACRO16:                  "MACRO16",
	KeyMACRO17:                  "MACRO17",
	KeyMACRO18:                  "MACRO18",
	KeyMACRO19:                  "MACRO19",
	KeyMACRO20:                  "MACRO20",
	KeyMACRO21:                  "MACRO21",
	KeyMACRO22:                  "MACRO22",
	KeyMACRO23:                  "MACRO23",
	KeyMACRO24:                  "MACRO24",
	KeyMACRO25:                  "MACRO25",
	KeyMACRO26:                  "MACRO26",
	KeyMACRO27:                  "MACRO27",
	KeyMACRO28:                  "MACRO28",
	KeyMACRO29:                  "MACRO29",
	KeyMACRO30:                  "MACRO30",
	KeyMACRO_RECORD_START:       "MACRO_RECORD_START",
	KeyMACRO_RECORD_STOP:        "MACRO_RECORD_STOP",
	KeyMACRO_PRESET_CYCLE:       "MACRO_PRESET_CYCLE",
	KeyMACRO_PRESET1:            "MACRO_PRESET1",
	KeyMACRO_PRESET2:            "MACRO_PRESET2",
	KeyMACRO_PRESET3:            "MACRO_PRESET3",
	KeyKBD_LCD_MENU1:            "KBD_LCD_MENU1",
	KeyKBD_LCD_MENU2:            "KBD_LCD_MENU2",
	KeyKBD_LCD_MENU3:            "KBD_LCD_MENU3",
	KeyKBD_LCD_MENU4:            "KBD_LCD_MENU4",
	KeyKBD_LCD_MENU5:            "KBD_LCD_MENU5",
	Btn0:                        "0",
	Btn1:                        "1",
	Btn2:                        "2",
	Btn3:                        "3",
	Btn4:                        "4",
	Btn5:                        "5",
	Btn6:                        "6",
	Btn7:                        "7",
	Btn8:                        "8",
	Btn9:                        "9",
	BtnLEFT:                     "LEFT",
	BtnRIGHT:                    "RIGHT",
	BtnMIDDLE:                   "MIDDLE",
	BtnSIDE:                     "SIDE",
	BtnEXTRA:                    "EXTRA",
	BtnFORWARD:                  "FORWARD",
	BtnBACK:                     "BACK",
	BtnTASK:                     "TASK",
	BtnTRIGGER:                  "TRIGGER",
	BtnTHUMB:                    "THUMB",
	BtnTHUMB2:                   "THUMB2",
	BtnTOP:                      "TOP",
	BtnTOP2:                     "TOP2",
	BtnPINKIE:                   "PINKIE",
	BtnBASE:                     "BASE",
	BtnBASE2:                    "BASE2",
	BtnBASE3:                    "BASE3",
	BtnBASE4:                    "BASE4",
	BtnBASE5:                    "BASE5",
	BtnBASE6:                    "BASE6",
	BtnDEAD:                     "DEAD",
	BtnSOUTH:                    "SOUTH",
	BtnEAST:                     "EAST",
	BtnC:                        "C",
	BtnNORTH:                    "NORTH",
	BtnWEST:                     "WEST",
	BtnZ:                        "Z",
	BtnTL:                       "TL",
	BtnTR:                       "TR",
	BtnTL2:                      "TL2",
	BtnTR2:                      "TR2",
	BtnSELECT:                   "SELECT",
	BtnSTART:                    "START",
	BtnMODE:                     "MODE",
	BtnTHUMBL:                   "THUMBL",
	BtnTHUMBR:                   "THUMBR",
	BtnTOOL_PEN:                 "TOOL_PEN",
	BtnTOOL_RUBBER:              "TOOL_RUBBER",
	BtnTOOL_BRUSH:               "TOOL_BRUSH",
	BtnTOOL_PENCIL:              "TOOL_PENCIL",
	BtnTOOL_AIRBRUSH:            "TOOL_AIRBRUSH",
	BtnTOOL_FINGER:              "TOOL_FINGER",
	BtnTOOL_MOUSE:               "TOOL_MOUSE",
	BtnTOOL_LENS:                "TOOL_LENS",
	BtnTOOL_QUINTTAP:            "TOOL_QUINTTAP",
	BtnSTYLUS3:                  "STYLUS3",
	BtnTOUCH:                    "TOUCH",
	BtnSTYLUS:                   "STYLUS",
	BtnSTYLUS2:                  "STYLUS2",
	BtnTOOL_DOUBLETAP:           "TOOL_DOUBLETAP",
	BtnTOOL_TRIPLETAP:           "TOOL_TRIPLETAP",
	BtnTOOL_QUADTAP:             "TOOL_QUADTAP",
	BtnGEAR_DOWN:                "GEAR_DOWN",
	BtnGEAR_UP:                  "GEAR_UP",
	BtnDPAD_UP:                  "DPAD_UP",
	BtnDPAD_DOWN:                "DPAD_DOWN",
	BtnDPAD_LEFT:                "DPAD_LEFT",
	BtnDPAD_RIGHT:               "DPAD_RIGHT",
	BtnTRIGGER_HAPPY1:           "TRIGGER_HAPPY1",
	BtnTRIGGER_HAPPY2:           "TRIGGER_HAPPY2",
	BtnTRIGGER_HAPPY3:           "TRIGGER_HAPPY3",
	BtnTRIGGER_HAPPY4:           "TRIGGER_HAPPY4",
	BtnTRIGGER_HAPPY5:           "TRIGGER_HAPPY5",
	BtnTRIGGER_HAPPY6:           "TRIGGER_HAPPY6",
	BtnTRIGGER_HAPPY7:           "TRIGGER_HAPPY7",
	BtnTRIGGER_HAPPY8:           "TRIGGER_HAPPY8",
	BtnTRIGGER_HAPPY9:           "TRIGGER_HAPPY9",
	BtnTRIGGER_HAPPY10:          "TRIGGER_HAPPY10",
	BtnTRIGGER_HAPPY11:          "TRIGGER_HAPPY11",
	BtnTRIGGER_HAPPY12:          "TRIGGER_HAPPY12",
	BtnTRIGGER_HAPPY13:          "TRIGGER_HAPPY13",
	BtnTRIGGER_HAPPY14:          "TRIGGER_HAPPY14",
	BtnTRIGGER_HAPPY15:          "TRIGGER_HAPPY15",
	BtnTRIGGER_HAPPY16:          "TRIGGER_HAPPY16",
	BtnTRIGGER_HAPPY17:          "TRIGGER_HAPPY17",
	BtnTRIGGER_HAPPY18:          "TRIGGER_HAPPY18",
	BtnTRIGGER_HAPPY19:          "TRIGGER_HAPPY19",
	BtnTRIGGER_HAPPY20:          "TRIGGER_HAPPY20",
	BtnTRIGGER_HAPPY21:          "TRIGGER_HAPPY21",
	BtnTRIGGER_HAPPY22:          "TRIGGER_HAPPY22",
	BtnTRIGGER_HAPPY23:          "TRIGGER_HAPPY23",
	BtnTRIGGER_HAPPY24:          "TRIGGER_HAPPY24",
	BtnTRIGGER_HAPPY25:          "TRIGGER_HAPPY25",
	BtnTRIGGER_HAPPY26:          "TRIGGER_HAPPY26",
	BtnTRIGGER_HAPPY27:          "TRIGGER_HAPPY27",
	BtnTRIGGER_HAPPY28:          "TRIGGER_HAPPY28",
	BtnTRIGGER_HAPPY29:          "TRIGGER_HAPPY29",
	BtnTRIGGER_HAPPY30:          "TRIGGER_HAPPY30",
	BtnTRIGGER_HAPPY31:          "TRIGGER_HAPPY31",
	BtnTRIGGER_HAPPY32:          "TRIGGER_HAPPY32",
	BtnTRIGGER_HAPPY33:          "TRIGGER_HAPPY33",
	BtnTRIGGER_HAPPY34:          "TRIGGER_HAPPY34",
	BtnTRIGGER_HAPPY35:          "TRIGGER_HAPPY35",
	BtnTRIGGER_HAPPY36:          "TRIGGER_HAPPY36",
	BtnTRIGGER_HAPPY37:          "TRIGGER_HAPPY37",
	BtnTRIGGER_HAPPY38:          "TRIGGER_HAPPY38",
	BtnTRIGGER_HAPPY39:          "TRIGGER_HAPPY39",
	BtnTRIGGER_HAPPY40:          "TRIGGER_HAPPY40",
}
