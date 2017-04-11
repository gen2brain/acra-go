package acra

import (
	"time"
)

// Report (Application Crash Reports for Android) struct
type Report struct {
	AndroidVersion   string  `json:"ANDROID_VERSION" form:"ANDROID_VERSION"`
	AppVersionCode   int     `json:"APP_VERSION_CODE" form:"APP_VERSION_CODE"`
	AppVersionName   float64 `json:"APP_VERSION_NAME" form:"APP_VERSION_NAME"`
	AvailableMemSize int     `json:"AVAILABLE_MEM_SIZE" form:"AVAILABLE_MEM_SIZE"`
	Brand            string  `json:"BRAND" form:"BRAND"`

	Build struct {
		Board        string `json:"BOARD" form:"BOARD"`
		Bootloader   string `json:"BOOTLOADER" form:"BOOTLOADER"`
		Brand        string `json:"BRAND" form:"BRAND"`
		CpuAbi       string `json:"CPU_ABI" form:"CPU_ABI"`
		CpuAbi2      string `json:"CPU_ABI2" form:"CPU_ABI2"`
		Device       string `json:"DEVICE" form:"DEVICE"`
		Display      string `json:"DISPLAY" form:"DISPLAY"`
		Fingerprint  string `json:"FINGERPRINT" form:"FINGERPRINT"`
		Hardware     string `json:"HARDWARE" form:"HARDWARE"`
		Host         string `json:"HOST" form:"HOST"`
		ID           string `json:"ID" form:"ID"`
		IsDebuggable bool   `json:"IS_DEBUGGABLE" form:"IS_DEBUGGABLE"`
		Manufacturer string `json:"MANUFACTURER" form:"MANUFACTURER"`
		Model        string `json:"MODEL" form:"MODEL"`
		Product      string `json:"PRODUCT" form:"PRODUCT"`
		Radio        string `json:"RADIO" form:"RADIO"`
		Serial       string `json:"SERIAL" form:"SERIAL"`
		Tags         string `json:"TAGS" form:"TAGS"`
		Time         int64  `json:"TIME" form:"TIME"`
		Type         string `json:"TYPE" form:"TYPE"`
		Unknown      string `json:"UNKNOWN" form:"UNKNOWN"`
		User         string `json:"USER" form:"USER"`

		Version struct {
			Codename        string `json:"CODENAME" form:"CODENAME"`
			Incremental     int    `json:"INCREMENTAL" form:"INCREMENTAL"`
			Release         string `json:"RELEASE" form:"RELEASE"`
			ResourcesSdkInt int    `json:"RESOURCES_SDK_INT" form:"RESOURCES_SDK_INT"`
			Sdk             int    `json:"SDK" form:"SDK"`
			SdkInt          int    `json:"SDK_INT" form:"SDK_INT"`
		} `json:"VERSION" form:"VERSION"`
	} `json:"BUILD" form:"BUILD"`

	CrashConfiguration struct {
		CompatScreenHeightDp        int    `json:"compatScreenHeightDp" form:"compatScreenHeightDp"`
		CompatScreenWidthDp         int    `json:"compatScreenWidthDp" form:"compatScreenWidthDp"`
		CompatSmallestScreenWidthDp int    `json:"compatSmallestScreenWidthDp" form:"compatSmallestScreenWidthDp"`
		DensityDpi                  int    `json:"densityDpi" form:"densityDpi"`
		FontScale                   string `json:"fontScale" form:"fontScale"`
		HardKeyboardHidden          string `json:"hardKeyboardHidden" form:"hardKeyboardHidden"`
		Keyboard                    string `json:"keyboard" form:"keyboard"`
		KeyboardHidden              string `json:"keyboardHidden" form:"keyboardHidden"`
		Locale                      string `json:"locale" form:"locale"`
		Mcc                         int    `json:"mcc" form:"mcc"`
		Mnc                         int    `json:"mnc" form:"mnc"`
		Navigation                  string `json:"navigation" form:"navigation"`
		NavigationHidden            string `json:"navigationHidden" form:"navigationHidden"`
		Orientation                 string `json:"orientation" form:"orientation"`
		ScreenHeightDp              int    `json:"screenHeightDp" form:"screenHeightDp"`
		ScreenLayout                string `json:"screenLayout" form:"screenLayout"`
		ScreenWidthDp               int    `json:"screenWidthDp" form:"screenWidthDp"`
		Seq                         int    `json:"seq" form:"seq"`
		SmallestScreenWidthDp       int    `json:"smallestScreenWidthDp" form:"smallestScreenWidthDp"`
		Touchscreen                 string `json:"touchscreen" form:"touchscreen"`
		UIMode                      string `json:"uiMode" form:"uiMode"`
		UserSetLocale               bool   `json:"userSetLocale" form:"userSetLocale"`
	} `json:"CRASH_CONFIGURATION" form:"CRASH_CONFIGURATION"`

	CustomData struct {
		Multiline string `json:"MULTILINE" form:"MULTILINE"`
		NoValue   string `json:"NO_VALUE" form:"NO_VALUE"`
	} `json:"CUSTOM_DATA" form:"CUSTOM_DATA"`

	DeviceFeatures struct {
		AndroidHardwareBluetooth                     bool   `json:"android.hardware.bluetooth" form:"android.hardware.bluetooth"`
		AndroidHardwareCamera                        bool   `json:"android.hardware.camera" form:"android.hardware.camera"`
		AndroidHardwareCameraAny                     bool   `json:"android.hardware.camera.any" form:"android.hardware.camera.any"`
		AndroidHardwareCameraAutofocus               bool   `json:"android.hardware.camera.autofocus" form:"android.hardware.camera.autofocus"`
		AndroidHardwareFaketouch                     bool   `json:"android.hardware.faketouch" form:"android.hardware.faketouch"`
		AndroidHardwareLocation                      bool   `json:"android.hardware.location" form:"android.hardware.location"`
		AndroidHardwareLocationNetwork               bool   `json:"android.hardware.location.network" form:"android.hardware.location.network"`
		AndroidHardwareMicrophone                    bool   `json:"android.hardware.microphone" form:"android.hardware.microphone"`
		AndroidHardwareScreenLandscape               bool   `json:"android.hardware.screen.landscape" form:"android.hardware.screen.landscape"`
		AndroidHardwareScreenPortrait                bool   `json:"android.hardware.screen.portrait" form:"android.hardware.screen.portrait"`
		AndroidHardwareSensorAccelerometer           bool   `json:"android.hardware.sensor.accelerometer" form:"android.hardware.sensor.accelerometer"`
		AndroidHardwareSensorCompass                 bool   `json:"android.hardware.sensor.compass" form:"android.hardware.sensor.compass"`
		AndroidHardwareTouchscreen                   bool   `json:"android.hardware.touchscreen" form:"android.hardware.touchscreen"`
		AndroidHardwareTouchscreenMultitouch         bool   `json:"android.hardware.touchscreen.multitouch" form:"android.hardware.touchscreen.multitouch"`
		AndroidHardwareTouchscreenMultitouchDistinct bool   `json:"android.hardware.touchscreen.multitouch.distinct" form:"android.hardware.touchscreen.multitouch.distinct"`
		AndroidHardwareTouchscreenMultitouchJazzhand bool   `json:"android.hardware.touchscreen.multitouch.jazzhand" form:"android.hardware.touchscreen.multitouch.jazzhand"`
		AndroidSoftwareLiveWallpaper                 bool   `json:"android.software.live_wallpaper" form:"android.software.live_wallpaper"`
		GlEsVersion                                  string `json:"glEsVersion" form:"glEsVersion"`
	} `json:"DEVICE_FEATURES" form:"DEVICE_FEATURES"`

	Display struct {
		Num0 struct {
			CurrentSizeRange struct {
				Largest  string `json:"largest" form:"largest"`
				Smallest string `json:"smallest" form:"smallest"`
			} `json:"currentSizeRange" form:"currentSizeRange"`

			Flags       string  `json:"flags" form:"flags"`
			GetRealSize string  `json:"getRealSize" form:"getRealSize"`
			GetSize     string  `json:"getSize" form:"getSize"`
			Height      int     `json:"height" form:"height"`
			IsValid     bool    `json:"isValid" form:"isValid"`
			Name        string  `json:"name" form:"name"`
			Orientation int     `json:"orientation" form:"orientation"`
			PixelFormat int     `json:"pixelFormat" form:"pixelFormat"`
			RectSize    string  `json:"rectSize" form:"rectSize"`
			RefreshRate float64 `json:"refreshRate" form:"refreshRate"`
			Rotation    string  `json:"rotation" form:"rotation"`
			Width       int     `json:"width" form:"width"`
		} `json:"0" form:"0"`
	} `json:"DISPLAY" form:"DISPLAY"`

	DumpsysMeminfo string `json:"DUMPSYS_MEMINFO" form:"DUMPSYS_MEMINFO"`

	Environment struct {
		GetDataDirectory                     string `json:"getDataDirectory" form:"getDataDirectory"`
		GetDownloadCacheDirectory            string `json:"getDownloadCacheDirectory" form:"getDownloadCacheDirectory"`
		GetEmulatedStorageObbSource          string `json:"getEmulatedStorageObbSource" form:"getEmulatedStorageObbSource"`
		GetExternalStorageAndroidDataDir     string `json:"getExternalStorageAndroidDataDir" form:"getExternalStorageAndroidDataDir"`
		GetExternalStorageDirectory          string `json:"getExternalStorageDirectory" form:"getExternalStorageDirectory"`
		GetExternalStorageState              string `json:"getExternalStorageState" form:"getExternalStorageState"`
		GetLegacyExternalStorageDirectory    string `json:"getLegacyExternalStorageDirectory" form:"getLegacyExternalStorageDirectory"`
		GetLegacyExternalStorageObbDirectory string `json:"getLegacyExternalStorageObbDirectory" form:"getLegacyExternalStorageObbDirectory"`
		GetMediaStorageDirectory             string `json:"getMediaStorageDirectory" form:"getMediaStorageDirectory"`
		GetRootDirectory                     string `json:"getRootDirectory" form:"getRootDirectory"`
		GetSecureDataDirectory               string `json:"getSecureDataDirectory" form:"getSecureDataDirectory"`
		GetSystemSecureDirectory             string `json:"getSystemSecureDirectory" form:"getSystemSecureDirectory"`
		IsEncryptedFilesystemEnabled         bool   `json:"isEncryptedFilesystemEnabled" form:"isEncryptedFilesystemEnabled"`
		IsExternalStorageEmulated            bool   `json:"isExternalStorageEmulated" form:"isExternalStorageEmulated"`
		IsExternalStorageRemovable           bool   `json:"isExternalStorageRemovable" form:"isExternalStorageRemovable"`
	} `json:"ENVIRONMENT" form:"ENVIRONMENT"`

	Filepath string `json:"FILE_PATH" form:"FILE_PATH"`

	InitialConfiguration struct {
		CompatScreenHeightDp        int    `json:"compatScreenHeightDp" form:"compatScreenHeightDp"`
		CompatScreenWidthDp         int    `json:"compatScreenWidthDp" form:"compatScreenWidthDp"`
		CompatSmallestScreenWidthDp int    `json:"compatSmallestScreenWidthDp" form:"compatSmallestScreenWidthDp"`
		DensityDpi                  int    `json:"densityDpi" form:"densityDpi"`
		FontScale                   string `json:"fontScale" form:"fontScale"`
		HardKeyboardHidden          string `json:"hardKeyboardHidden" form:"hardKeyboardHidden"`
		Keyboard                    string `json:"keyboard" form:"keyboard"`
		KeyboardHidden              string `json:"keyboardHidden" form:"keyboardHidden"`
		Locale                      string `json:"locale" form:"locale"`
		Mcc                         int    `json:"mcc" form:"mcc"`
		Mnc                         int    `json:"mnc" form:"mnc"`
		Navigation                  string `json:"navigation" form:"navigation"`
		NavigationHidden            string `json:"navigationHidden" form:"navigationHidden"`
		Orientation                 string `json:"orientation" form:"orientation"`
		ScreenHeightDp              int    `json:"screenHeightDp" form:"screenHeightDp"`
		ScreenLayout                string `json:"screenLayout" form:"screenLayout"`
		ScreenWidthDp               int    `json:"screenWidthDp" form:"screenWidthDp"`
		Seq                         int    `json:"seq" form:"seq"`
		SmallestScreenWidthDp       int    `json:"smallestScreenWidthDp" form:"smallestScreenWidthDp"`
		Touchscreen                 string `json:"touchscreen" form:"touchscreen"`
		UIMode                      string `json:"uiMode" form:"uiMode"`
		UserSetLocale               bool   `json:"userSetLocale" form:"userSetLocale"`
	} `json:"INITIAL_CONFIGURATION" form:"INITIAL_CONFIGURATION"`

	InstallationID string `json:"INSTALLATION_ID" form:"INSTALLATION_ID"`
	IsSilent       bool   `json:"IS_SILENT" form:"IS_SILENT"`
	Logcat         string `json:"LOGCAT" form:"LOGCAT"`
	PackageName    string `json:"PACKAGE_NAME" form:"PACKAGE_NAME"`
	PhoneModel     string `json:"PHONE_MODEL" form:"PHONE_MODEL"`
	Product        string `json:"PRODUCT" form:"PRODUCT"`
	ReportID       string `json:"REPORT_ID" form:"REPORT_ID"`

	SettingsGlobal struct {
		AirplaneModeOn                      int    `json:"AIRPLANE_MODE_ON" form:"AIRPLANE_MODE_ON"`
		AirplaneModeRadios                  string `json:"AIRPLANE_MODE_RADIOS" form:"AIRPLANE_MODE_RADIOS"`
		AirplaneModeToggleableRadios        string `json:"AIRPLANE_MODE_TOGGLEABLE_RADIOS" form:"AIRPLANE_MODE_TOGGLEABLE_RADIOS"`
		AssistedGpsEnabled                  int    `json:"ASSISTED_GPS_ENABLED" form:"ASSISTED_GPS_ENABLED"`
		AudioSafeVolumeState                int    `json:"AUDIO_SAFE_VOLUME_STATE" form:"AUDIO_SAFE_VOLUME_STATE"`
		AutoTime                            int    `json:"AUTO_TIME" form:"AUTO_TIME"`
		AutoTimeZone                        int    `json:"AUTO_TIME_ZONE" form:"AUTO_TIME_ZONE"`
		BluetoothOn                         int    `json:"BLUETOOTH_ON" form:"BLUETOOTH_ON"`
		CallAutoRetry                       int    `json:"CALL_AUTO_RETRY" form:"CALL_AUTO_RETRY"`
		CarDockSound                        string `json:"CAR_DOCK_SOUND" form:"CAR_DOCK_SOUND"`
		CarUndockSound                      string `json:"CAR_UNDOCK_SOUND" form:"CAR_UNDOCK_SOUND"`
		CdmaCellBroadcastSms                int    `json:"CDMA_CELL_BROADCAST_SMS" form:"CDMA_CELL_BROADCAST_SMS"`
		DataRoaming                         int    `json:"DATA_ROAMING" form:"DATA_ROAMING"`
		DefaultInstallLocation              int    `json:"DEFAULT_INSTALL_LOCATION" form:"DEFAULT_INSTALL_LOCATION"`
		DeskDockSound                       string `json:"DESK_DOCK_SOUND" form:"DESK_DOCK_SOUND"`
		DeskUndockSound                     string `json:"DESK_UNDOCK_SOUND" form:"DESK_UNDOCK_SOUND"`
		DeviceProvisioned                   int    `json:"DEVICE_PROVISIONED" form:"DEVICE_PROVISIONED"`
		DockSoundsEnabled                   int    `json:"DOCK_SOUNDS_ENABLED" form:"DOCK_SOUNDS_ENABLED"`
		EmergencyTone                       int    `json:"EMERGENCY_TONE" form:"EMERGENCY_TONE"`
		InstallNonMarketApps                int    `json:"INSTALL_NON_MARKET_APPS" form:"INSTALL_NON_MARKET_APPS"`
		LockSound                           string `json:"LOCK_SOUND" form:"LOCK_SOUND"`
		LowBatterySound                     string `json:"LOW_BATTERY_SOUND" form:"LOW_BATTERY_SOUND"`
		MobileData                          int    `json:"MOBILE_DATA" form:"MOBILE_DATA"`
		ModeRinger                          int    `json:"MODE_RINGER" form:"MODE_RINGER"`
		NetstatsEnabled                     int    `json:"NETSTATS_ENABLED" form:"NETSTATS_ENABLED"`
		NetworkPreference                   int    `json:"NETWORK_PREFERENCE" form:"NETWORK_PREFERENCE"`
		PackageVerifierEnable               int    `json:"PACKAGE_VERIFIER_ENABLE" form:"PACKAGE_VERIFIER_ENABLE"`
		PowerSoundsEnabled                  int    `json:"POWER_SOUNDS_ENABLED" form:"POWER_SOUNDS_ENABLED"`
		PreferredNetworkMode                int    `json:"PREFERRED_NETWORK_MODE" form:"PREFERRED_NETWORK_MODE"`
		SetInstallLocation                  int    `json:"SET_INSTALL_LOCATION" form:"SET_INSTALL_LOCATION"`
		StayOnWhilePluggedIn                int    `json:"STAY_ON_WHILE_PLUGGED_IN" form:"STAY_ON_WHILE_PLUGGED_IN"`
		ThrottleResetDay                    int    `json:"THROTTLE_RESET_DAY" form:"THROTTLE_RESET_DAY"`
		UnlockSound                         string `json:"UNLOCK_SOUND" form:"UNLOCK_SOUND"`
		UsbMassStorageEnabled               int    `json:"USB_MASS_STORAGE_ENABLED" form:"USB_MASS_STORAGE_ENABLED"`
		WifiCountryCode                     string `json:"WIFI_COUNTRY_CODE" form:"WIFI_COUNTRY_CODE"`
		WifiDisplayOn                       int    `json:"WIFI_DISPLAY_ON" form:"WIFI_DISPLAY_ON"`
		WifiMaxDhcpRetryCount               int    `json:"WIFI_MAX_DHCP_RETRY_COUNT" form:"WIFI_MAX_DHCP_RETRY_COUNT"`
		WifiNetworksAvailableNotificationOn int    `json:"WIFI_NETWORKS_AVAILABLE_NOTIFICATION_ON" form:"WIFI_NETWORKS_AVAILABLE_NOTIFICATION_ON"`
		WifiOn                              int    `json:"WIFI_ON" form:"WIFI_ON"`
		WifiSleepPolicy                     int    `json:"WIFI_SLEEP_POLICY" form:"WIFI_SLEEP_POLICY"`
		WifiWatchdogOn                      int    `json:"WIFI_WATCHDOG_ON" form:"WIFI_WATCHDOG_ON"`
		WirelessChargingStartedSound        string `json:"WIRELESS_CHARGING_STARTED_SOUND" form:"WIRELESS_CHARGING_STARTED_SOUND"`
	} `json:"SETTINGS_GLOBAL" form:"SETTINGS_GLOBAL"`

	SettingsSecure struct {
		AccessibilityDisplayMagnificationAutoUpdate int    `json:"ACCESSIBILITY_DISPLAY_MAGNIFICATION_AUTO_UPDATE" form:"ACCESSIBILITY_DISPLAY_MAGNIFICATION_AUTO_UPDATE"`
		AccessibilityDisplayMagnificationEnabled    int    `json:"ACCESSIBILITY_DISPLAY_MAGNIFICATION_ENABLED" form:"ACCESSIBILITY_DISPLAY_MAGNIFICATION_ENABLED"`
		AccessibilityDisplayMagnificationScale      string `json:"ACCESSIBILITY_DISPLAY_MAGNIFICATION_SCALE" form:"ACCESSIBILITY_DISPLAY_MAGNIFICATION_SCALE"`
		AccessibilityScreenreaderUrl                string `json:"ACCESSIBILITY_SCREEN_READER_URL" form:"ACCESSIBILITY_SCREEN_READER_URL"`
		AccessibilityScriptInjection                int    `json:"ACCESSIBILITY_SCRIPT_INJECTION" form:"ACCESSIBILITY_SCRIPT_INJECTION"`
		AccessibilitySpeakPassword                  int    `json:"ACCESSIBILITY_SPEAK_PASSWORD" form:"ACCESSIBILITY_SPEAK_PASSWORD"`
		AccessibilityWebContentKeyBindings          string `json:"ACCESSIBILITY_WEB_CONTENT_KEY_BINDINGS" form:"ACCESSIBILITY_WEB_CONTENT_KEY_BINDINGS"`
		AllowMockLocation                           int    `json:"ALLOW_MOCK_LOCATION" form:"ALLOW_MOCK_LOCATION"`
		AndroidID                                   string `json:"ANDROID_ID" form:"ANDROID_ID"`
		BackupEnabled                               int    `json:"BACKUP_ENABLED" form:"BACKUP_ENABLED"`
		BackupTransport                             string `json:"BACKUP_TRANSPORT" form:"BACKUP_TRANSPORT"`
		DefaultInputMethod                          string `json:"DEFAULT_INPUT_METHOD" form:"DEFAULT_INPUT_METHOD"`
		EnabledInputMethods                         string `json:"ENABLED_INPUT_METHODS" form:"ENABLED_INPUT_METHODS"`
		InputMethodsSubtypeHistory                  string `json:"INPUT_METHODS_SUBTYPE_HISTORY" form:"INPUT_METHODS_SUBTYPE_HISTORY"`
		LocationProvidersAllowed                    string `json:"LOCATION_PROVIDERS_ALLOWED" form:"LOCATION_PROVIDERS_ALLOWED"`
		LockPatternEnabled                          int    `json:"LOCK_PATTERN_ENABLED" form:"LOCK_PATTERN_ENABLED"`
		LockPatternVisible                          int    `json:"LOCK_PATTERN_VISIBLE" form:"LOCK_PATTERN_VISIBLE"`
		LongPressTimeout                            int    `json:"LONG_PRESS_TIMEOUT" form:"LONG_PRESS_TIMEOUT"`
		MountPlayNotificationSnd                    int    `json:"MOUNT_PLAY_NOTIFICATION_SND" form:"MOUNT_PLAY_NOTIFICATION_SND"`
		MountUmsAutostart                           int    `json:"MOUNT_UMS_AUTOSTART" form:"MOUNT_UMS_AUTOSTART"`
		MountUmsNotifyEnabled                       int    `json:"MOUNT_UMS_NOTIFY_ENABLED" form:"MOUNT_UMS_NOTIFY_ENABLED"`
		MountUmsPrompt                              int    `json:"MOUNT_UMS_PROMPT" form:"MOUNT_UMS_PROMPT"`
		ScreensaverActivateOnDock                   int    `json:"SCREENSAVER_ACTIVATE_ON_DOCK" form:"SCREENSAVER_ACTIVATE_ON_DOCK"`
		ScreensaverActivateOnSleep                  int    `json:"SCREENSAVER_ACTIVATE_ON_SLEEP" form:"SCREENSAVER_ACTIVATE_ON_SLEEP"`
		ScreensaverComponents                       string `json:"SCREENSAVER_COMPONENTS" form:"SCREENSAVER_COMPONENTS"`
		ScreensaverDefaultComponent                 string `json:"SCREENSAVER_DEFAULT_COMPONENT" form:"SCREENSAVER_DEFAULT_COMPONENT"`
		ScreensaverEnabled                          int    `json:"SCREENSAVER_ENABLED" form:"SCREENSAVER_ENABLED"`
		SelectedInputMethodSubtype                  string `json:"SELECTED_INPUT_METHOD_SUBTYPE" form:"SELECTED_INPUT_METHOD_SUBTYPE"`
		SelectedSpellChecker                        string `json:"SELECTED_SPELL_CHECKER" form:"SELECTED_SPELL_CHECKER"`
		SelectedSpellCheckerSubtype                 int    `json:"SELECTED_SPELL_CHECKER_SUBTYPE" form:"SELECTED_SPELL_CHECKER_SUBTYPE"`
		TouchExplorationEnabled                     int    `json:"TOUCH_EXPLORATION_ENABLED" form:"TOUCH_EXPLORATION_ENABLED"`
		UserSetupComplete                           int    `json:"USER_SETUP_COMPLETE" form:"USER_SETUP_COMPLETE"`
	} `json:"SETTINGS_SECURE" form:"SETTINGS_SECURE"`

	SettingsSystem struct {
		AccelerometerRotation     int    `json:"ACCELEROMETER_ROTATION" form:"ACCELEROMETER_ROTATION"`
		DtmfToneTypeWhenDialing   int    `json:"DTMF_TONE_TYPE_WHEN_DIALING" form:"DTMF_TONE_TYPE_WHEN_DIALING"`
		DtmfToneWhenDialing       int    `json:"DTMF_TONE_WHEN_DIALING" form:"DTMF_TONE_WHEN_DIALING"`
		HapticFeedbackEnabled     int    `json:"HAPTIC_FEEDBACK_ENABLED" form:"HAPTIC_FEEDBACK_ENABLED"`
		HearingAid                int    `json:"HEARING_AID" form:"HEARING_AID"`
		LockscreenSoundsEnabled   int    `json:"LOCKSCREEN_SOUNDS_ENABLED" form:"LOCKSCREEN_SOUNDS_ENABLED"`
		MediaButtonReceiver       string `json:"MEDIA_BUTTON_RECEIVER" form:"MEDIA_BUTTON_RECEIVER"`
		ModeRingerStreamsAffected int    `json:"MODE_RINGER_STREAMS_AFFECTED" form:"MODE_RINGER_STREAMS_AFFECTED"`
		MuteStreamsAffected       int    `json:"MUTE_STREAMS_AFFECTED" form:"MUTE_STREAMS_AFFECTED"`
		NextAlarmFormatted        string `json:"NEXT_ALARM_FORMATTED" form:"NEXT_ALARM_FORMATTED"`
		NotificationLightPulse    int    `json:"NOTIFICATION_LIGHT_PULSE" form:"NOTIFICATION_LIGHT_PULSE"`
		PointerSpeed              int    `json:"POINTER_SPEED" form:"POINTER_SPEED"`
		ScreenBrightness          int    `json:"SCREEN_BRIGHTNESS" form:"SCREEN_BRIGHTNESS"`
		ScreenBrightnessMode      int    `json:"SCREEN_BRIGHTNESS_MODE" form:"SCREEN_BRIGHTNESS_MODE"`
		ScreenOffTimeout          int    `json:"SCREEN_OFF_TIMEOUT" form:"SCREEN_OFF_TIMEOUT"`
		SoundEffectsEnabled       int    `json:"SOUND_EFFECTS_ENABLED" form:"SOUND_EFFECTS_ENABLED"`
		TtyMode                   int    `json:"TTY_MODE" form:"TTY_MODE"`
		VibrateWhenRinging        int    `json:"VIBRATE_WHEN_RINGING" form:"VIBRATE_WHEN_RINGING"`
		VolumeAlarm               int    `json:"VOLUME_ALARM" form:"VOLUME_ALARM"`
		VolumeBluetoothSco        int    `json:"VOLUME_BLUETOOTH_SCO" form:"VOLUME_BLUETOOTH_SCO"`
		VolumeMusic               int    `json:"VOLUME_MUSIC" form:"VOLUME_MUSIC"`
		VolumeNotification        int    `json:"VOLUME_NOTIFICATION" form:"VOLUME_NOTIFICATION"`
		VolumeRing                int    `json:"VOLUME_RING" form:"VOLUME_RING"`
		VolumeSystem              int    `json:"VOLUME_SYSTEM" form:"VOLUME_SYSTEM"`
		VolumeVoice               int    `json:"VOLUME_VOICE" form:"VOLUME_VOICE"`
	} `json:"SETTINGS_SYSTEM" form:"SETTINGS_SYSTEM"`

	SharedPreferences struct {
		Default struct {
			Acra struct {
				LastVersionNr int `json:"lastVersionNr" form:"lastVersionNr"`
			} `json:"acra" form:"acra"`
		} `json:"default" form:"default"`
	} `json:"SHARED_PREFERENCES" form:"SHARED_PREFERENCES"`

	StackTrace       string    `json:"STACK_TRACE" form:"STACK_TRACE"`
	TotalMemSize     int       `json:"TOTAL_MEM_SIZE" form:"TOTAL_MEM_SIZE"`
	UserAppStartDate time.Time `json:"USER_APP_START_DATE" form:"USER_APP_START_DATE"`
	UserCrashDate    time.Time `json:"USER_CRASH_DATE" form:"USER_CRASH_DATE"`
	UserEmail        string    `json:"USER_EMAIL" form:"USER_EMAIL"`
}
