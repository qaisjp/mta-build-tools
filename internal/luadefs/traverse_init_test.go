package luadefs_test

import (
	"os"
	"path"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
)

func dummyFS() billy.Filesystem {
	f := memfs.New()

	// Create all the files
	for filepath, content := range dummyFiles {
		f.MkdirAll(path.Dir(filepath), os.ModeDir)

		file, err := f.Create(filepath)
		Must(err)

		_, err = file.Write([]byte(content))
		Must(err)

		Must(file.Close())
	}

	return f
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

var dummyServerFuncNames = []string{
	// Server/mods/deathmatch/logic/lua/CLuaManager.cpp
	"addEvent", "addEventHandler", "removeEventHandler", "getEventHandlers", "triggerEvent", "triggerClientEvent", "cancelEvent", "wasEventCancelled", "getCancelReason", "triggerLatentClientEvent", "getLatentEventHandles", "getLatentEventStatus", "cancelLatentEvent", "addDebugHook", "removeDebugHook", "createExplosion", "getBodyPartName", "getClothesByTypeIndex", "getTypeIndexFromClothes", "getClothesTypeName", "getWeaponNameFromID", "getWeaponIDFromName", "getWeaponProperty", "getOriginalWeaponProperty", "setWeaponProperty", "setWeaponAmmo", "getSlotFromWeapon", "createWeapon", "fireWeapon", "setWeaponState", "getWeaponState", "setWeaponTarget", "getWeaponTarget", "setWeaponOwner", "getWeaponOwner", "setWeaponFlags", "getWeaponFlags", "setWeaponFiringRate", "getWeaponFiringRate", "resetWeaponFiringRate", "getWeaponAmmo", "getWeaponClipAmmo", "setWeaponClipAmmo", "addCommandHandler", "removeCommandHandler", "executeCommandHandler", "getCommandHandlers", "getMaxPlayers", "setMaxPlayers", "outputChatBox", "outputConsole", "outputDebugString", "outputServerLog", "getServerName", "getServerHttpPort", "getServerPassword", "setServerPassword", "getServerConfigSetting", "clearChatBox", "getRootElement", "loadMapData", "saveMapData", "getGameType", "getMapName", "setGameType", "setMapName", "getRuleValue", "setRuleValue", "removeRuleValue", "getPerformanceStats", "callRemote", "fetchRemote", "resetMapInfo", "getServerPort", "get", "set", "getVersion", "getNetworkUsageData", "getNetworkStats", "getLoadedModules", "getModuleInfo", "setDevelopmentMode", "getDevelopmentMode", "getPlayerSkin", "setPlayerSkin", "getVehicleModel", "setVehicleModel", "getObjectModel", "setObjectModel", "getVehicleID", "getVehicleIDFromName", "getVehicleNameFromID", "getPlayerWeaponSlot", "getPlayerArmor", "getPlayerRotation", "isPlayerChoking", "isPlayerDead", "isPlayerDucked", "getPlayerStat", "getPlayerTarget", "getPlayerClothes", "doesPlayerHaveJetPack", "isPlayerInWater", "isPedInWater", "isPlayerOnGround", "getPlayerFightingStyle", "getPlayerGravity", "getPlayerContactElement", "setPlayerArmor", "setPlayerWeaponSlot", "killPlayer", "setPlayerRotation", "setPlayerStat", "addPlayerClothes", "removePlayerClothes", "givePlayerJetPack", "removePlayerJetPack", "setPlayerFightingStyle", "setPlayerGravity", "setPlayerChoking", "warpPlayerIntoVehicle", "removePlayerFromVehicle", "getPlayerOccupiedVehicle", "getPlayerOccupiedVehicleSeat", "isPlayerInVehicle", "getClientName", "getClientIP", "getClientAccount", "setClientName", "getPlayerWeapon", "getPlayerTotalAmmo", "getPlayerAmmoInClip", "getPedSkin", "setPedSkin", "xmlNodeGetSubNodes", "xmlCreateSubNode", "xmlFindSubNode", "attachElementToElement", "detachElementFromElement", "doesPedHaveJetPack", "givePedJetPack", "removePedJetPack",

	// Server/mods/deathmatch/logic/luadefs/CLuaCameraDefs.cpp
	"getCameraMatrix", "getCameraTarget", "getCameraInterior", "setCameraMatrix", "setCameraTarget", "setCameraInterior", "fadeCamera",
}

var dummyClientFuncNames = []string{
	// Client/mods/deathmatch/logic/lua/CLuaManager.cpp
	"getPlayerRotation", "canPlayerBeKnockedOffBike", "getPlayerContactElement", "isPlayerInVehicle", "doesPlayerHaveJetPack", "isPlayerInWater", "isPedInWater", "isPedOnFire", "setPedOnFire", "isPlayerOnGround", "getPlayerTask", "getPlayerSimplestTask", "isPlayerDoingTask", "getPlayerTarget", "getPlayerTargetStart", "getPlayerTargetEnd", "getPlayerTargetRange", "getPlayerTargetCollision", "getPlayerWeaponSlot", "getPlayerWeapon", "getPlayerAmmoInClip", "getPlayerTotalAmmo", "getPedWeaponMuzzlePosition", "getPlayerOccupiedVehicle", "getPlayerArmor", "getPlayerSkin", "isPlayerChoking", "isPlayerDucked", "getPlayerStat", "setPlayerWeaponSlot", "setPlayerSkin", "setPlayerRotation", "setPlayerCanBeKnockedOffBike", "setVehicleModel", "getVehicleModel", "getPedSkin", "setPedSkin", "getObjectRotation", "setObjectRotation", "getVehicleIDFromName", "getVehicleID", "getVehicleRotation", "getVehicleNameFromID", "setVehicleRotation", "attachElementToElement", "detachElementFromElement", "xmlFindSubNode", "xmlNodeGetSubNodes", "xmlNodeFindSubNode", "xmlCreateSubNode", "xmlNodeFindChild", "isPlayerDead", "guiEditSetCaratIndex", "guiMemoSetCaratIndex", "setControlState", "getControlState", "addEvent", "addEventHandler", "removeEventHandler", "getEventHandlers", "triggerEvent", "triggerServerEvent", "cancelEvent", "wasEventCancelled", "triggerLatentServerEvent", "getLatentEventHandles", "getLatentEventStatus", "cancelLatentEvent", "outputConsole", "outputChatBox", "showChat", "isChatVisible", "outputDebugString", "setClipboard", "setWindowFlashing", "clearChatBox", "createTrayNotification", "isTrayNotificationEnabled", "getBodyPartName", "getClothesByTypeIndex", "getTypeIndexFromClothes", "getClothesTypeName", "createExplosion", "getCursorPosition", "setCursorPosition", "isCursorShowing", "showCursor", "getCursorAlpha", "setCursorAlpha", "getValidPedModels", "downloadFile", "getTime", "getGroundPosition", "processLineOfSight", "getWorldFromScreenPosition", "getScreenFromWorldPosition", "getWeather", "getZoneName", "getGravity", "getGameSpeed", "getMinuteDuration", "getWaveHeight", "getGaragePosition", "getGarageSize", "getGarageBoundingBox", "getBlurLevel", "getTrafficLightState", "areTrafficLightsLocked", "getSkyGradient", "getHeatHaze", "getJetpackMaxHeight", "getWindVelocity", "getInteriorSoundsEnabled", "getInteriorFurnitureEnabled", "getFarClipDistance", "getNearClipDistance", "getVehiclesLODDistance", "getPedsLODDistance", "getFogDistance", "getSunColor", "getSunSize", "getAircraftMaxHeight", "getAircraftMaxVelocity", "getOcclusionsEnabled", "getCloudsEnabled", "getRainLevel", "getMoonSize", "getFPSLimit", "getBirdsEnabled", "isPedTargetingMarkerEnabled", "isLineOfSightClear", "isWorldSpecialPropertyEnabled", "isGarageOpen", "setTime", "setSkyGradient", "setHeatHaze", "setWeather", "setWeatherBlended", "setGravity", "setGameSpeed", "setWaveHeight", "setMinuteDuration", "setGarageOpen", "setWorldSpecialPropertyEnabled", "setBlurLevel", "setJetpackMaxHeight", "setCloudsEnabled", "setTrafficLightState", "setTrafficLightsLocked", "setWindVelocity", "setInteriorSoundsEnabled", "setInteriorFurnitureEnabled", "setRainLevel", "setFarClipDistance", "setNearClipDistance", "setVehiclesLODDistance", "setPedsLODDistance", "setFogDistance", "setSunColor", "setSunSize", "setAircraftMaxHeight", "setAircraftMaxVelocity", "setOcclusionsEnabled", "setBirdsEnabled", "setPedTargetingMarkerEnabled", "setMoonSize", "setFPSLimit", "removeWorldModel", "restoreAllWorldModels", "restoreWorldModel", "createSWATRope", "resetSkyGradient", "resetHeatHaze", "resetWindVelocity", "resetRainLevel", "resetFarClipDistance", "resetNearClipDistance", "resetVehiclesLODDistance", "resetPedsLODDistance", "resetFogDistance", "resetSunColor", "resetSunSize", "resetMoonSize", "bindKey", "unbindKey", "getKeyState", "getAnalogControlState", "setAnalogControlState", "isControlEnabled", "getBoundKeys", "getFunctionsBoundToKey", "getKeyBoundToFunction", "getCommandsBoundToKey", "getKeyBoundToCommand", "toggleControl", "toggleAllControls", "addCommandHandler", "removeCommandHandler", "executeCommandHandler", "getCommandHandlers", "getNetworkUsageData", "getNetworkStats", "getPerformanceStats", "setDevelopmentMode", "getDevelopmentMode", "addDebugHook", "removeDebugHook", "fetchRemote", "getVersion", "getLocalization", "getKeyboardLayout", "isVoiceEnabled",

	// Client/mods/deathmatch/logic/luadefs/CLuaBlipDefs.cpp
	"createBlip", "createBlipAttachedTo", "getBlipIcon", "getBlipSize", "getBlipColor", "getBlipOrdering", "getBlipVisibleDistance", "setBlipIcon", "setBlipSize", "setBlipColor", "setBlipOrdering", "setBlipVisibleDistance",

	// Client/mods/deathmatch/logic/luadefs/CLuaClassDefs.cpp
	// <none>
}

var dummyFiles = map[string]string{
	"Server/mods/deathmatch/logic/lua/CLuaManager.cpp": `/*****************************************************************************
*
*  PROJECT:     Multi Theft Auto v1.0
*  LICENSE:     See LICENSE in the top level directory
*  FILE:        mods/deathmatch/logic/lua/CLuaManager.cpp
*  PURPOSE:     Lua virtual machine manager class
*
*  Multi Theft Auto is available from http://www.multitheftauto.com/
*
*****************************************************************************/

#include "StdInc.h"

extern CGame* g_pGame;

CLuaManager::CLuaManager(CObjectManager* pObjectManager, CPlayerManager* pPlayerManager, CVehicleManager* pVehicleManager, CBlipManager* pBlipManager,
						CRadarAreaManager* pRadarAreaManager, CRegisteredCommands* pRegisteredCommands, CMapManager* pMapManager, CEvents* pEvents)
{
	m_pObjectManager = pObjectManager;
	m_pPlayerManager = pPlayerManager;
	m_pVehicleManager = pVehicleManager;
	m_pBlipManager = pBlipManager;
	m_pRadarAreaManager = pRadarAreaManager;
	m_pRegisteredCommands = pRegisteredCommands;
	m_pMapManager = pMapManager;
	m_pEvents = pEvents;

	// Create our lua dynamic module manager
	m_pLuaModuleManager = new CLuaModuleManager(this);
	m_pLuaModuleManager->SetScriptDebugging(g_pGame->GetScriptDebugging());

	// Load our C Functions into Lua and hook callback
	LoadCFunctions();
	lua_registerPreCallHook(CLuaDefs::CanUseFunction);
	lua_registerUndumpHook(CLuaMain::OnUndump);

#ifdef MTA_DEBUG
	// Check rounding in case json is updated
	json_object* obj = json_object_new_double(5.1);
	SString      strResult = json_object_to_json_string_ext(obj, JSON_C_TO_STRING_PLAIN);
	assert(strResult == "5.1");
	json_object_put(obj);
#endif
}

CLuaManager::~CLuaManager()
{
	CLuaCFunctions::RemoveAllFunctions();
	list<CLuaMain*>::iterator iter;
	for (iter = m_virtualMachines.begin(); iter != m_virtualMachines.end(); ++iter)
	{
		delete (*iter);
	}

	// Destroy the module manager
	delete m_pLuaModuleManager;
}

CLuaMain* CLuaManager::CreateVirtualMachine(CResource* pResourceOwner, bool bEnableOOP)
{
	// Create it and add it to the list over VM's
	CLuaMain* pLuaMain = new CLuaMain(this, m_pObjectManager, m_pPlayerManager, m_pVehicleManager, m_pBlipManager, m_pRadarAreaManager, m_pMapManager,
										pResourceOwner, bEnableOOP);
	m_virtualMachines.push_back(pLuaMain);
	pLuaMain->InitVM();

	m_pLuaModuleManager->RegisterFunctions(pLuaMain->GetVirtualMachine());

	return pLuaMain;
}

bool CLuaManager::RemoveVirtualMachine(CLuaMain* pLuaMain)
{
	if (pLuaMain)
	{
		// Remove all events registered by it and all commands added
		m_pEvents->RemoveAllEvents(pLuaMain);
		m_pRegisteredCommands->CleanUpForVM(pLuaMain);

		// Delete it unless it is already
		if (!pLuaMain->BeingDeleted())
		{
			delete pLuaMain;
		}

		// Remove it from our list
		m_virtualMachines.remove(pLuaMain);
		return true;
	}

	return false;
}

void CLuaManager::OnLuaMainOpenVM(CLuaMain* pLuaMain, lua_State* luaVM)
{
	MapSet(m_VirtualMachineMap, pLuaMain->GetVirtualMachine(), pLuaMain);
}

void CLuaManager::OnLuaMainCloseVM(CLuaMain* pLuaMain, lua_State* luaVM)
{
	MapRemove(m_VirtualMachineMap, pLuaMain->GetVirtualMachine());
}

void CLuaManager::DoPulse()
{
	list<CLuaMain*>::iterator iter;
	for (iter = m_virtualMachines.begin(); iter != m_virtualMachines.end(); ++iter)
	{
		(*iter)->DoPulse();
	}
	m_pLuaModuleManager->DoPulse();
}

CLuaMain* CLuaManager::GetVirtualMachine(lua_State* luaVM)
{
	if (!luaVM)
		return NULL;

	// Grab the main virtual state because the one we've got passed might be a coroutine state
	// and only the main state is in our list.
	lua_State* main = lua_getmainstate(luaVM);
	if (main)
	{
		luaVM = main;
	}

	// Find a matching VM in our map
	CLuaMain* pLuaMain = MapFindRef(m_VirtualMachineMap, luaVM);
	if (pLuaMain)
		return pLuaMain;

	// Find a matching VM in our list
	list<CLuaMain*>::const_iterator iter = m_virtualMachines.begin();
	for (; iter != m_virtualMachines.end(); ++iter)
	{
		if (luaVM == (*iter)->GetVirtualMachine())
		{
			dassert(0);            // Why not in map?
			return *iter;
		}
	}

	// Doesn't exist
	return NULL;
}

// Return resource associated with a lua state
CResource* CLuaManager::GetVirtualMachineResource(lua_State* luaVM)
{
	CLuaMain* pLuaMain = GetVirtualMachine(luaVM);
	if (pLuaMain)
		return pLuaMain->GetResource();
	return NULL;
}

void CLuaManager::LoadCFunctions()
{
	std::map<const char*, lua_CFunction> functions{
		{"addEvent", CLuaFunctionDefs::AddEvent},
		{"addEventHandler", CLuaFunctionDefs::AddEventHandler},
		{"removeEventHandler", CLuaFunctionDefs::RemoveEventHandler},
		{"getEventHandlers", CLuaFunctionDefs::GetEventHandlers},
		{"triggerEvent", CLuaFunctionDefs::TriggerEvent},
		{"triggerClientEvent", CLuaFunctionDefs::TriggerClientEvent},
		{"cancelEvent", CLuaFunctionDefs::CancelEvent},
		{"wasEventCancelled", CLuaFunctionDefs::WasEventCancelled},
		{"getCancelReason", CLuaFunctionDefs::GetCancelReason},
		{"triggerLatentClientEvent", CLuaFunctionDefs::TriggerLatentClientEvent},
		{"getLatentEventHandles", CLuaFunctionDefs::GetLatentEventHandles},
		{"getLatentEventStatus", CLuaFunctionDefs::GetLatentEventStatus},
		{"cancelLatentEvent", CLuaFunctionDefs::CancelLatentEvent},
		{"addDebugHook", CLuaFunctionDefs::AddDebugHook},
		{"removeDebugHook", CLuaFunctionDefs::RemoveDebugHook},

		// Explosion create funcs
		{"createExplosion", CLuaFunctionDefs::CreateExplosion},

		// Fire create funcs
		// CLuaCFunctions::AddFunction ( "createFire", CLuaFunctionDefinitions::CreateFire );

		// Path(node) funcs
		// CLuaCFunctions::AddFunction ( "createNode", CLuaFunctionDefinitions::CreateNode );

		// Ped body funcs?
		{"getBodyPartName", CLuaFunctionDefs::GetBodyPartName},
		{"getClothesByTypeIndex", CLuaFunctionDefs::GetClothesByTypeIndex},
		{"getTypeIndexFromClothes", CLuaFunctionDefs::GetTypeIndexFromClothes},
		{"getClothesTypeName", CLuaFunctionDefs::GetClothesTypeName},

		// Weapon funcs
		{"getWeaponNameFromID", CLuaFunctionDefs::GetWeaponNameFromID},
		{"getWeaponIDFromName", CLuaFunctionDefs::GetWeaponIDFromName},
		{"getWeaponProperty", CLuaFunctionDefs::GetWeaponProperty},
		{"getOriginalWeaponProperty", CLuaFunctionDefs::GetOriginalWeaponProperty},
		{"setWeaponProperty", CLuaFunctionDefs::SetWeaponProperty},
		{"setWeaponAmmo", CLuaFunctionDefs::SetWeaponAmmo},
		{"getSlotFromWeapon", CLuaFunctionDefs::GetSlotFromWeapon},

	#if MTASA_VERSION_TYPE < VERSION_TYPE_RELEASE
		{"createWeapon", CLuaFunctionDefs::CreateWeapon},
		{"fireWeapon", CLuaFunctionDefs::FireWeapon},
		{"setWeaponState", CLuaFunctionDefs::SetWeaponState},
		{"getWeaponState", CLuaFunctionDefs::GetWeaponState},
		{"setWeaponTarget", CLuaFunctionDefs::SetWeaponTarget},
		{"getWeaponTarget", CLuaFunctionDefs::GetWeaponTarget},
		{"setWeaponOwner", CLuaFunctionDefs::SetWeaponOwner},
		{"getWeaponOwner", CLuaFunctionDefs::GetWeaponOwner},
		{"setWeaponFlags", CLuaFunctionDefs::SetWeaponFlags},
		{"getWeaponFlags", CLuaFunctionDefs::GetWeaponFlags},
		{"setWeaponFiringRate", CLuaFunctionDefs::SetWeaponFiringRate},
		{"getWeaponFiringRate", CLuaFunctionDefs::GetWeaponFiringRate},
		{"resetWeaponFiringRate", CLuaFunctionDefs::ResetWeaponFiringRate},
		{"getWeaponAmmo", CLuaFunctionDefs::GetWeaponAmmo},
		{"getWeaponClipAmmo", CLuaFunctionDefs::GetWeaponClipAmmo},
		{"setWeaponClipAmmo", CLuaFunctionDefs::SetWeaponClipAmmo},
	#endif

		// Console funcs
		{"addCommandHandler", CLuaFunctionDefs::AddCommandHandler},
		{"removeCommandHandler", CLuaFunctionDefs::RemoveCommandHandler},
		{"executeCommandHandler", CLuaFunctionDefs::ExecuteCommandHandler},
		{"getCommandHandlers", CLuaFunctionDefs::GetCommandHandlers},

		// Server standard funcs
		{"getMaxPlayers", CLuaFunctionDefs::GetMaxPlayers},
		{"setMaxPlayers", CLuaFunctionDefs::SetMaxPlayers},
		{"outputChatBox", CLuaFunctionDefs::OutputChatBox},
		{"outputConsole", CLuaFunctionDefs::OutputConsole},
		{"outputDebugString", CLuaFunctionDefs::OutputDebugString},
		{"outputServerLog", CLuaFunctionDefs::OutputServerLog},
		{"getServerName", CLuaFunctionDefs::GetServerName},
		{"getServerHttpPort", CLuaFunctionDefs::GetServerHttpPort},
		{"getServerPassword", CLuaFunctionDefs::GetServerPassword},
		{"setServerPassword", CLuaFunctionDefs::SetServerPassword},
		{"getServerConfigSetting", CLuaFunctionDefs::GetServerConfigSetting},
		{"clearChatBox", CLuaFunctionDefs::ClearChatBox},

		// Loaded map funcs
		{"getRootElement", CLuaFunctionDefs::GetRootElement},
		{"loadMapData", CLuaFunctionDefs::LoadMapData},
		{"saveMapData", CLuaFunctionDefs::SaveMapData},

		// All-Seeing Eye Functions
		{"getGameType", CLuaFunctionDefs::GetGameType},
		{"getMapName", CLuaFunctionDefs::GetMapName},
		{"setGameType", CLuaFunctionDefs::SetGameType},
		{"setMapName", CLuaFunctionDefs::SetMapName},
		{"getRuleValue", CLuaFunctionDefs::GetRuleValue},
		{"setRuleValue", CLuaFunctionDefs::SetRuleValue},
		{"removeRuleValue", CLuaFunctionDefs::RemoveRuleValue},

		// Registry functions
		{"getPerformanceStats", CLuaFunctionDefs::GetPerformanceStats},

		// Admin functions
		/*
		CLuaCFunctions::AddFunction ( "aexec", CLuaFunctionDefinitions::Aexec },
		CLuaCFunctions::AddFunction ( "kickPlayer", CLuaFunctionDefinitions::KickPlayer },
		CLuaCFunctions::AddFunction ( "banPlayer", CLuaFunctionDefinitions::BanPlayer },
		CLuaCFunctions::AddFunction ( "banPlayerIP", CLuaFunctionDefinitions::BanPlayerIP },
		CLuaCFunctions::AddFunction ( "setPlayerMuted", CLuaFunctionDefinitions::SetPlayerMuted },

		CLuaCFunctions::AddFunction ( "addAccount", CLuaFunctionDefinitions::AddAccount },
		CLuaCFunctions::AddFunction ( "delAccount", CLuaFunctionDefinitions::DelAccount },
		CLuaCFunctions::AddFunction ( "setAccountPassword", CLuaFunctionDefinitions::SetAccountPassword },
		*/

		{"callRemote", CLuaFunctionDefs::CallRemote},
		{"fetchRemote", CLuaFunctionDefs::FetchRemote},

		// Misc funcs
		{"resetMapInfo", CLuaFunctionDefs::ResetMapInfo},
		{"getServerPort", CLuaFunctionDefs::GetServerPort},

		// Settings registry funcs
		{"get", CLuaFunctionDefs::Get},
		{"set", CLuaFunctionDefs::Set},

		// Utility
		{"getVersion", CLuaFunctionDefs::GetVersion},
		{"getNetworkUsageData", CLuaFunctionDefs::GetNetworkUsageData},
		{"getNetworkStats", CLuaFunctionDefs::GetNetworkStats},
		{"getLoadedModules", CLuaFunctionDefs::GetModules},
		{"getModuleInfo", CLuaFunctionDefs::GetModuleInfo},

		{"setDevelopmentMode", CLuaFunctionDefs::SetDevelopmentMode},
		{"getDevelopmentMode", CLuaFunctionDefs::GetDevelopmentMode},

		// Backward compat functions at the end, so the new function name is used in ACL

		// ** BACKWARDS COMPATIBILITY FUNCS. SHOULD BE REMOVED BEFORE FINAL RELEASE! **
		{"getPlayerSkin", CLuaElementDefs::getElementModel},
		{"setPlayerSkin", CLuaElementDefs::setElementModel},
		{"getVehicleModel", CLuaElementDefs::getElementModel},
		{"setVehicleModel", CLuaElementDefs::setElementModel},
		{"getObjectModel", CLuaElementDefs::getElementModel},
		{"setObjectModel", CLuaElementDefs::setElementModel},
		{"getVehicleID", CLuaElementDefs::getElementModel},
		{"getVehicleIDFromName", CLuaVehicleDefs::GetVehicleModelFromName},
		{"getVehicleNameFromID", CLuaVehicleDefs::GetVehicleNameFromModel},
		{"getPlayerWeaponSlot", CLuaPedDefs::GetPedWeaponSlot},
		{"getPlayerArmor", CLuaPedDefs::GetPedArmor},
		{"getPlayerRotation", CLuaPedDefs::GetPedRotation},
		{"isPlayerChoking", CLuaPedDefs::IsPedChoking},
		{"isPlayerDead", CLuaPedDefs::IsPedDead},
		{"isPlayerDucked", CLuaPedDefs::IsPedDucked},
		{"getPlayerStat", CLuaPedDefs::GetPedStat},
		{"getPlayerTarget", CLuaPedDefs::GetPedTarget},
		{"getPlayerClothes", CLuaPedDefs::GetPedClothes},
		{"doesPlayerHaveJetPack", CLuaPedDefs::DoesPedHaveJetPack},
		{"isPlayerInWater", CLuaElementDefs::isElementInWater},
		{"isPedInWater", CLuaElementDefs::isElementInWater},
		{"isPlayerOnGround", CLuaPedDefs::IsPedOnGround},
		{"getPlayerFightingStyle", CLuaPedDefs::GetPedFightingStyle},
		{"getPlayerGravity", CLuaPedDefs::GetPedGravity},
		{"getPlayerContactElement", CLuaPedDefs::GetPedContactElement},
		{"setPlayerArmor", CLuaPedDefs::SetPedArmor},
		{"setPlayerWeaponSlot", CLuaPedDefs::SetPedWeaponSlot},
		{"killPlayer", CLuaPedDefs::KillPed},
		{"setPlayerRotation", CLuaPedDefs::SetPedRotation},
		{"setPlayerStat", CLuaPedDefs::SetPedStat},
		{"addPlayerClothes", CLuaPedDefs::AddPedClothes},
		{"removePlayerClothes", CLuaPedDefs::RemovePedClothes},
		{"givePlayerJetPack", CLuaPedDefs::GivePedJetPack},
		{"removePlayerJetPack", CLuaPedDefs::RemovePedJetPack},
		{"setPlayerFightingStyle", CLuaPedDefs::SetPedFightingStyle},
		{"setPlayerGravity", CLuaPedDefs::SetPedGravity},
		{"setPlayerChoking", CLuaPedDefs::SetPedChoking},
		{"warpPlayerIntoVehicle", CLuaPedDefs::WarpPedIntoVehicle},
		{"removePlayerFromVehicle", CLuaPedDefs::RemovePedFromVehicle},
		{"getPlayerOccupiedVehicle", CLuaPedDefs::GetPedOccupiedVehicle},
		{"getPlayerOccupiedVehicleSeat", CLuaPedDefs::GetPedOccupiedVehicleSeat},
		{"isPlayerInVehicle", CLuaPedDefs::IsPedInVehicle},
		{"getClientName", CLuaPlayerDefs::GetPlayerName},
		{"getClientIP", CLuaPlayerDefs::GetPlayerIP},
		{"getClientAccount", CLuaPlayerDefs::GetPlayerAccount},
		{"setClientName", CLuaPlayerDefs::SetPlayerName},
		{"getPlayerWeapon", CLuaPedDefs::GetPedWeapon},
		{"getPlayerTotalAmmo", CLuaPedDefs::GetPedTotalAmmo},
		{"getPlayerAmmoInClip", CLuaPedDefs::GetPedAmmoInClip},
		{"getPedSkin", CLuaElementDefs::getElementModel},
		{"setPedSkin", CLuaElementDefs::setElementModel},
		{"xmlNodeGetSubNodes", CLuaXMLDefs::xmlNodeGetChildren},
		{"xmlCreateSubNode", CLuaXMLDefs::xmlCreateChild},
		{"xmlFindSubNode", CLuaXMLDefs::xmlNodeFindChild},
		{"attachElementToElement", CLuaElementDefs::attachElements},
		{"detachElementFromElement", CLuaElementDefs::detachElements},

		// Deprecated since 1.5.5-9.13846
		{"doesPedHaveJetPack", CLuaPedDefs::DoesPedHaveJetPack},
		{"givePedJetPack", CLuaPedDefs::GivePedJetPack},
		{"removePedJetPack", CLuaPedDefs::RemovePedJetPack}
		// ** END OF BACKWARDS COMPATIBILITY FUNCS. **
	};

	// Add all functions
	for (const auto& pair : functions)
	{
		CLuaCFunctions::AddFunction(pair.first, pair.second);
	}

	// Restricted functions
	CLuaCFunctions::AddFunction("setServerConfigSetting", CLuaFunctionDefs::SetServerConfigSetting, true);
	CLuaCFunctions::AddFunction("shutdown", CLuaFunctionDefs::shutdown, true);

	// Load the functions from our classes
	CLuaACLDefs::LoadFunctions();
	CLuaAccountDefs::LoadFunctions();
	CLuaBanDefs::LoadFunctions();
	CLuaBlipDefs::LoadFunctions();
	CLuaCameraDefs::LoadFunctions();
	CLuaColShapeDefs::LoadFunctions();
	CLuaDatabaseDefs::LoadFunctions();
	CLuaElementDefs::LoadFunctions();
	CLuaHandlingDefs::LoadFunctions();
	CLuaMarkerDefs::LoadFunctions();
	CLuaObjectDefs::LoadFunctions();
	CLuaPedDefs::LoadFunctions();
	CLuaPickupDefs::LoadFunctions();
	CLuaPlayerDefs::LoadFunctions();
	CLuaRadarAreaDefs::LoadFunctions();
	CLuaResourceDefs::LoadFunctions();
	CLuaShared::LoadFunctions();
	CLuaTeamDefs::LoadFunctions();
	CLuaTextDefs::LoadFunctions();
	CLuaTimerDefs::LoadFunctions();
	CLuaVehicleDefs::LoadFunctions();
	CLuaVoiceDefs::LoadFunctions();
	CLuaWaterDefs::LoadFunctions();
	CLuaWorldDefs::LoadFunctions();
	CLuaXMLDefs::LoadFunctions();
}`,
	"Server/mods/deathmatch/logic/luadefs/CLuaCameraDefs.cpp": `/*****************************************************************************
*
*  PROJECT:     Multi Theft Auto v1.0
*  LICENSE:     See LICENSE in the top level directory
*  FILE:        mods/deathmatch/logic/luadefs/CLuaCameraDefs.cpp
*  PURPOSE:     Lua camera function definitions class
*
*  Multi Theft Auto is available from http://www.multitheftauto.com/
*
*****************************************************************************/

#include "StdInc.h"

void CLuaCameraDefs::LoadFunctions()
{
	std::map<const char*, lua_CFunction> functions{
		// Get functions
		{"getCameraMatrix", getCameraMatrix},
		{"getCameraTarget", getCameraTarget},
		{"getCameraInterior", getCameraInterior},

		// Set functions
		{"setCameraMatrix", setCameraMatrix},
		{"setCameraTarget", setCameraTarget},
		{"setCameraInterior", setCameraInterior},
		{"fadeCamera", fadeCamera},
	};

	// Add functions
	for (const auto& pair : functions)
	{
		CLuaCFunctions::AddFunction(pair.first, pair.second);
	}
}

int CLuaCameraDefs::getCameraMatrix(lua_State* luaVM)
{
	//  float cameraX, float cameraY, float cameraZ, float targetX, float targetY, float targetZ, float roll, float fov getCameraMatrix ( player thePlayer )
	CPlayer* pPlayer;

	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pPlayer);

	if (!argStream.HasErrors())
	{
		CVector vecPosition, vecLookAt;
		float   fRoll, fFOV;
		if (CStaticFunctionDefinitions::GetCameraMatrix(pPlayer, vecPosition, vecLookAt, fRoll, fFOV))
		{
			lua_pushnumber(luaVM, vecPosition.fX);
			lua_pushnumber(luaVM, vecPosition.fY);
			lua_pushnumber(luaVM, vecPosition.fZ);
			lua_pushnumber(luaVM, vecLookAt.fX);
			lua_pushnumber(luaVM, vecLookAt.fY);
			lua_pushnumber(luaVM, vecLookAt.fZ);
			lua_pushnumber(luaVM, fRoll);
			lua_pushnumber(luaVM, fFOV);
			return 8;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaCameraDefs::OOP_getCameraMatrix(lua_State* luaVM)
{
	// Needs further attention before adding
	CPlayer* pPlayer;

	CScriptArgReader argStream(luaVM);

	argStream.ReadUserData(pPlayer);

	if (!argStream.HasErrors())
	{
		CMatrix matrix;
		// pPlayer->GetCamera ()->GetMatrix ( matrix );

		lua_pushmatrix(luaVM, matrix);
		return 1;
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaCameraDefs::getCameraTarget(lua_State* luaVM)
{
	//  element getCameraTarget ( player thePlayer )
	CPlayer* pPlayer;

	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pPlayer);

	if (!argStream.HasErrors())
	{
		CElement* pTarget = CStaticFunctionDefinitions::GetCameraTarget(pPlayer);
		if (pTarget)
		{
			lua_pushelement(luaVM, pTarget);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaCameraDefs::getCameraInterior(lua_State* luaVM)
{
	//  int getCameraInterior ( player thePlayer )
	CPlayer* pPlayer;

	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pPlayer);

	if (!argStream.HasErrors())
	{
		unsigned char ucInterior;
		if (CStaticFunctionDefinitions::GetCameraInterior(pPlayer, ucInterior))
		{
			lua_pushnumber(luaVM, ucInterior);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaCameraDefs::setCameraMatrix(lua_State* luaVM)
{
	//  bool setCameraMatrix ( player thePlayer, float positionX, float positionY, float positionZ [, float lookAtX, float lookAtY, float lookAtZ, float roll =
	//  0, float fov = 70 ] )
	CElement* pPlayer;
	CVector   vecPosition;
	CVector   vecLookAt;
	float     fRoll;
	float     fFOV;

	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pPlayer);

	if (argStream.NextIsUserDataOfType<CLuaMatrix>())
	{
		CLuaMatrix* pMatrix;
		argStream.ReadUserData(pMatrix);

		vecPosition = pMatrix->GetPosition();
		vecLookAt = pMatrix->GetRotation();
	}
	else
	{
		argStream.ReadVector3D(vecPosition);
		argStream.ReadVector3D(vecLookAt, CVector());
	}

	argStream.ReadNumber(fRoll, 0.0f);
	argStream.ReadNumber(fFOV, 70.0f);

	if (!argStream.HasErrors())
	{
		if (fFOV <= 0.0f || fFOV >= 180.0f)
			fFOV = 70.0f;

		if (CStaticFunctionDefinitions::SetCameraMatrix(pPlayer, vecPosition, &vecLookAt, fRoll, fFOV))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaCameraDefs::setCameraTarget(lua_State* luaVM)
{
	//  bool setCameraTarget ( player thePlayer [, element target = nil ] )
	CElement* pPlayer;
	CElement* pTarget;

	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pPlayer);
	argStream.ReadUserData(pTarget, NULL);

	if (!argStream.HasErrors())
	{
		if (CStaticFunctionDefinitions::SetCameraTarget(pPlayer, pTarget))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaCameraDefs::setCameraInterior(lua_State* luaVM)
{
	//  bool setCameraInterior ( player thePlayer, int interior )
	CElement*     pElement;
	unsigned char ucInterior;

	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pElement);
	argStream.ReadNumber(ucInterior);

	if (!argStream.HasErrors())
	{
		if (CStaticFunctionDefinitions::SetCameraInterior(pElement, ucInterior))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaCameraDefs::fadeCamera(lua_State* luaVM)
{
	//  bool fadeCamera ( player thePlayer, bool fadeIn, [ float timeToFade = 1.0, int red = 0, int green = 0, int blue = 0 ] )
	CElement*     pPlayer;
	bool          bFadeIn;
	float         fFadeTime;
	unsigned char ucRed;
	unsigned char ucGreen;
	unsigned char ucBlue;

	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pPlayer);
	argStream.ReadBool(bFadeIn);
	argStream.ReadNumber(fFadeTime, 1.0f);
	argStream.ReadNumber(ucRed, 0);
	argStream.ReadNumber(ucGreen, 0);
	argStream.ReadNumber(ucBlue, 0);

	if (!argStream.HasErrors())
	{
		if (CStaticFunctionDefinitions::FadeCamera(pPlayer, bFadeIn, fFadeTime, ucRed, ucGreen, ucBlue))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}`,
	"Client/mods/deathmatch/logic/lua/CLuaManager.cpp": `/*****************************************************************************
*
*  PROJECT:     Multi Theft Auto v1.0
*  LICENSE:     See LICENSE in the top level directory
*  FILE:        mods/shared_logic/logic/lua/CLuaManager.cpp
*  PURPOSE:     Lua virtual machine manager class
*
*  Multi Theft Auto is available from http://www.multitheftauto.com/
*
*****************************************************************************/

#include "StdInc.h"
#include "../luadefs/CLuaFireDefs.h"

using std::list;

static int DummyPreCall(lua_CFunction f, lua_State* L)
{
	// Always allow functions
	return true;
}

CLuaManager::CLuaManager(CClientGame* pClientGame)
{
	m_pEvents = pClientGame->GetEvents();
	m_pGUIManager = pClientGame->GetGUIManager();
	m_pRegisteredCommands = pClientGame->GetRegisteredCommands();

	// Ensure lua was compiled with apichecks
	#ifdef NDEBUG
		#error "NDEBUG should not be defined"
	#endif
	assert(luaX_is_apicheck_enabled());

	// Load the C functions
	LoadCFunctions();
	lua_registerPreCallHook(CLuaDefs::CanUseFunction);
	lua_registerUndumpHook(CLuaMain::OnUndump);
}

CLuaManager::~CLuaManager()
{
	// Delete all the VM's
	list<CLuaMain*>::iterator iter;
	for (iter = m_virtualMachines.begin(); iter != m_virtualMachines.end(); iter++)
	{
		delete (*iter);
	}

	// Close and remove LVM from memory
	ProcessPendingDeleteList();

	// Clear the C functions
	CLuaCFunctions::RemoveAllFunctions();
}

CLuaMain* CLuaManager::CreateVirtualMachine(CResource* pResourceOwner, bool bEnableOOP)
{
	// Create it and add it to the list over VM's
	CLuaMain* pLuaMain = new CLuaMain(this, pResourceOwner, bEnableOOP);
	m_virtualMachines.push_back(pLuaMain);
	pLuaMain->InitVM();
	return pLuaMain;
}

bool CLuaManager::RemoveVirtualMachine(CLuaMain* pLuaMain)
{
	if (pLuaMain)
	{
		// Remove all events registered by it
		m_pEvents->RemoveAllEvents(pLuaMain);
		m_pRegisteredCommands->CleanUpForVM(pLuaMain);

		// Remove it from our list
		m_virtualMachines.remove(pLuaMain);

		// Delete it unless it is already
		if (!pLuaMain->BeingDeleted())
		{
			delete pLuaMain;
		}

		return true;
	}

	return false;
}

void CLuaManager::OnLuaMainOpenVM(CLuaMain* pLuaMain, lua_State* luaVM)
{
	MapSet(m_VirtualMachineMap, pLuaMain->GetVirtualMachine(), pLuaMain);
}

void CLuaManager::OnLuaMainCloseVM(CLuaMain* pLuaMain, lua_State* luaVM)
{
	MapRemove(m_VirtualMachineMap, pLuaMain->GetVirtualMachine());
}

void CLuaManager::ProcessPendingDeleteList()
{
	while (!m_PendingDeleteList.empty())
	{
		lua_State* luaVM = m_PendingDeleteList.front();
		m_PendingDeleteList.pop_front();
		CLuaFunctionRef::RemoveLuaFunctionRefsForVM(luaVM);
		lua_close(luaVM);
	}
}

void CLuaManager::DoPulse()
{
	list<CLuaMain*>::iterator iter = m_virtualMachines.begin();
	for (; iter != m_virtualMachines.end(); ++iter)
	{
		(*iter)->DoPulse();
	}
}

CLuaMain* CLuaManager::GetVirtualMachine(lua_State* luaVM)
{
	if (!luaVM)
		return NULL;

	// Grab the main virtual state because the one we've got passed might be a coroutine state
	// and only the main state is in our list.
	lua_State* main = lua_getmainstate(luaVM);
	if (main)
	{
		luaVM = main;
	}

	// Find a matching VM in our map
	CLuaMain* pLuaMain = MapFindRef(m_VirtualMachineMap, luaVM);
	if (pLuaMain)
		return pLuaMain;

	// Find a matching VM in our list (should not be needed)
	list<CLuaMain*>::const_iterator iter = m_virtualMachines.begin();
	for (; iter != m_virtualMachines.end(); iter++)
	{
		if (luaVM == (*iter)->GetVirtualMachine())
		{
			dassert(0);            // Why not in map?
			return *iter;
		}
	}

	// Doesn't exist
	return NULL;
}

void CLuaManager::LoadCFunctions()
{
	std::map<const char*, lua_CFunction> functions{
		// ** BACKWARDS COMPATIBILITY FUNCS. SHOULD BE REMOVED BEFORE FINAL RELEASE! **
		{"getPlayerRotation", CLuaPedDefs::GetPedRotation},
		{"canPlayerBeKnockedOffBike", CLuaPedDefs::CanPedBeKnockedOffBike},
		{"getPlayerContactElement", CLuaPedDefs::GetPedContactElement},
		{"isPlayerInVehicle", CLuaPedDefs::IsPedInVehicle},
		{"doesPlayerHaveJetPack", CLuaPedDefs::DoesPedHaveJetPack},
		{"isPlayerInWater", CLuaElementDefs::IsElementInWater},
		{"isPedInWater", CLuaElementDefs::IsElementInWater},
		{"isPedOnFire", CLuaPedDefs::IsPedOnFire},
		{"setPedOnFire", CLuaPedDefs::SetPedOnFire},
		{"isPlayerOnGround", CLuaPedDefs::IsPedOnGround},
		{"getPlayerTask", CLuaPedDefs::GetPedTask},
		{"getPlayerSimplestTask", CLuaPedDefs::GetPedSimplestTask},
		{"isPlayerDoingTask", CLuaPedDefs::IsPedDoingTask},
		{"getPlayerTarget", CLuaPedDefs::GetPedTarget},
		{"getPlayerTargetStart", CLuaPedDefs::GetPedTargetStart},
		{"getPlayerTargetEnd", CLuaPedDefs::GetPedTargetEnd},
		{"getPlayerTargetRange", CLuaPedDefs::GetPedTargetRange},
		{"getPlayerTargetCollision", CLuaPedDefs::GetPedTargetCollision},
		{"getPlayerWeaponSlot", CLuaPedDefs::GetPedWeaponSlot},
		{"getPlayerWeapon", CLuaPedDefs::GetPedWeapon},
		{"getPlayerAmmoInClip", CLuaPedDefs::GetPedAmmoInClip},
		{"getPlayerTotalAmmo", CLuaPedDefs::GetPedTotalAmmo},
		{"getPedWeaponMuzzlePosition", CLuaPedDefs::GetPedWeaponMuzzlePosition},
		{"getPlayerOccupiedVehicle", CLuaPedDefs::GetPedOccupiedVehicle},
		{"getPlayerArmor", CLuaPedDefs::GetPedArmor},
		{"getPlayerSkin", CLuaElementDefs::GetElementModel},
		{"isPlayerChoking", CLuaPedDefs::IsPedChoking},
		{"isPlayerDucked", CLuaPedDefs::IsPedDucked},
		{"getPlayerStat", CLuaPedDefs::GetPedStat},
		{"setPlayerWeaponSlot", CLuaPedDefs::SetPedWeaponSlot},
		{"setPlayerSkin", CLuaElementDefs::SetElementModel},
		{"setPlayerRotation", CLuaPedDefs::SetPedRotation},
		{"setPlayerCanBeKnockedOffBike", CLuaPedDefs::SetPedCanBeKnockedOffBike},
		{"setVehicleModel", CLuaElementDefs::SetElementModel},
		{"getVehicleModel", CLuaElementDefs::GetElementModel},
		{"getPedSkin", CLuaElementDefs::GetElementModel},
		{"setPedSkin", CLuaElementDefs::SetElementModel},
		{"getObjectRotation", CLuaElementDefs::GetElementRotation},
		{"setObjectRotation", CLuaElementDefs::SetElementRotation},
		{"getVehicleIDFromName", CLuaVehicleDefs::GetVehicleModelFromName},
		{"getVehicleID", CLuaElementDefs::GetElementModel},
		{"getVehicleRotation", CLuaElementDefs::GetElementRotation},
		{"getVehicleNameFromID", CLuaVehicleDefs::GetVehicleNameFromModel},
		{"setVehicleRotation", CLuaElementDefs::SetElementRotation},
		{"attachElementToElement", CLuaElementDefs::AttachElements},
		{"detachElementFromElement", CLuaElementDefs::DetachElements},
		{"xmlFindSubNode", CLuaXMLDefs::xmlNodeFindChild},
		{"xmlNodeGetSubNodes", CLuaXMLDefs::xmlNodeGetChildren},
		{"xmlNodeFindSubNode", CLuaXMLDefs::xmlNodeFindChild},
		{"xmlCreateSubNode", CLuaXMLDefs::xmlCreateChild},
		{"xmlNodeFindChild", CLuaXMLDefs::xmlNodeFindChild},
		{"isPlayerDead", CLuaPedDefs::IsPedDead},
		{"guiEditSetCaratIndex", CLuaGUIDefs::GUIEditSetCaretIndex},
		{"guiMemoSetCaratIndex", CLuaGUIDefs::GUIMemoSetCaretIndex},
		{"setControlState", CLuaPedDefs::SetPedControlState},
		{"getControlState", CLuaPedDefs::GetPedControlState},
		// ** END OF BACKWARDS COMPATIBILITY FUNCS. **

		// Event funcs
		{"addEvent", CLuaFunctionDefs::AddEvent},
		{"addEventHandler", CLuaFunctionDefs::AddEventHandler},
		{"removeEventHandler", CLuaFunctionDefs::RemoveEventHandler},
		{"getEventHandlers", CLuaFunctionDefs::GetEventHandlers},
		{"triggerEvent", CLuaFunctionDefs::TriggerEvent},
		{"triggerServerEvent", CLuaFunctionDefs::TriggerServerEvent},
		{"cancelEvent", CLuaFunctionDefs::CancelEvent},
		{"wasEventCancelled", CLuaFunctionDefs::WasEventCancelled},
		{"triggerLatentServerEvent", CLuaFunctionDefs::TriggerLatentServerEvent},
		{"getLatentEventHandles", CLuaFunctionDefs::GetLatentEventHandles},
		{"getLatentEventStatus", CLuaFunctionDefs::GetLatentEventStatus},
		{"cancelLatentEvent", CLuaFunctionDefs::CancelLatentEvent},

		// Output funcs
		{"outputConsole", CLuaFunctionDefs::OutputConsole},
		{"outputChatBox", CLuaFunctionDefs::OutputChatBox},
		{"showChat", CLuaFunctionDefs::ShowChat},
		{"isChatVisible", CLuaFunctionDefs::IsChatVisible},
		{"outputDebugString", CLuaFunctionDefs::OutputClientDebugString},
		{"setClipboard", CLuaFunctionDefs::SetClipboard},
		{"setWindowFlashing", CLuaFunctionDefs::SetWindowFlashing},
		{"clearChatBox", CLuaFunctionDefs::ClearChatBox},

		// Notification funcs
		{"createTrayNotification", CLuaFunctionDefs::CreateTrayNotification},
		{"isTrayNotificationEnabled", CLuaFunctionDefs::IsTrayNotificationEnabled},

		// Clothes and body functions
		{"getBodyPartName", CLuaFunctionDefs::GetBodyPartName},
		{"getClothesByTypeIndex", CLuaFunctionDefs::GetClothesByTypeIndex},
		{"getTypeIndexFromClothes", CLuaFunctionDefs::GetTypeIndexFromClothes},
		{"getClothesTypeName", CLuaFunctionDefs::GetClothesTypeName},

		// Explosion funcs
		{"createExplosion", CLuaFunctionDefs::CreateExplosion},

		// Cursor funcs
		{"getCursorPosition", CLuaFunctionDefs::GetCursorPosition},
		{"setCursorPosition", CLuaFunctionDefs::SetCursorPosition},
		{"isCursorShowing", CLuaFunctionDefs::IsCursorShowing},
		{"showCursor", CLuaFunctionDefs::ShowCursor},
		{"getCursorAlpha", CLuaFunctionDefs::GetCursorAlpha},
		{"setCursorAlpha", CLuaFunctionDefs::SetCursorAlpha},

		// Util functions
		{"getValidPedModels", CLuaFunctionDefs::GetValidPedModels},
		{"downloadFile", CLuaFunctionDefs::DownloadFile},

		// World get functions
		{"getTime", CLuaFunctionDefs::GetTime_},
		{"getGroundPosition", CLuaFunctionDefs::GetGroundPosition},
		{"processLineOfSight", CLuaFunctionDefs::ProcessLineOfSight},
		{"getWorldFromScreenPosition", CLuaFunctionDefs::GetWorldFromScreenPosition},
		{"getScreenFromWorldPosition", CLuaFunctionDefs::GetScreenFromWorldPosition},
		{"getWeather", CLuaFunctionDefs::GetWeather},
		{"getZoneName", CLuaFunctionDefs::GetZoneName},
		{"getGravity", CLuaFunctionDefs::GetGravity},
		{"getGameSpeed", CLuaFunctionDefs::GetGameSpeed},
		{"getMinuteDuration", CLuaFunctionDefs::GetMinuteDuration},
		{"getWaveHeight", CLuaFunctionDefs::GetWaveHeight},
		{"getGaragePosition", CLuaFunctionDefs::GetGaragePosition},
		{"getGarageSize", CLuaFunctionDefs::GetGarageSize},
		{"getGarageBoundingBox", CLuaFunctionDefs::GetGarageBoundingBox},
		{"getBlurLevel", CLuaFunctionDefs::GetBlurLevel},
		{"getTrafficLightState", CLuaFunctionDefs::GetTrafficLightState},
		{"areTrafficLightsLocked", CLuaFunctionDefs::AreTrafficLightsLocked},
		{"getSkyGradient", CLuaFunctionDefs::GetSkyGradient},
		{"getHeatHaze", CLuaFunctionDefs::GetHeatHaze},
		{"getJetpackMaxHeight", CLuaFunctionDefs::GetJetpackMaxHeight},
		{"getWindVelocity", CLuaFunctionDefs::GetWindVelocity},
		{"getInteriorSoundsEnabled", CLuaFunctionDefs::GetInteriorSoundsEnabled},
		{"getInteriorFurnitureEnabled", CLuaFunctionDefs::GetInteriorFurnitureEnabled},
		{"getFarClipDistance", CLuaFunctionDefs::GetFarClipDistance},
		{"getNearClipDistance", CLuaFunctionDefs::GetNearClipDistance},
		{"getVehiclesLODDistance", CLuaFunctionDefs::GetVehiclesLODDistance},
		{"getPedsLODDistance", CLuaFunctionDefs::GetPedsLODDistance},
		{"getFogDistance", CLuaFunctionDefs::GetFogDistance},
		{"getSunColor", CLuaFunctionDefs::GetSunColor},
		{"getSunSize", CLuaFunctionDefs::GetSunSize},
		{"getAircraftMaxHeight", CLuaFunctionDefs::GetAircraftMaxHeight},
		{"getAircraftMaxVelocity", CLuaFunctionDefs::GetAircraftMaxVelocity},
		{"getOcclusionsEnabled", CLuaFunctionDefs::GetOcclusionsEnabled},
		{"getCloudsEnabled", CLuaFunctionDefs::GetCloudsEnabled},
		{"getRainLevel", CLuaFunctionDefs::GetRainLevel},
		{"getMoonSize", CLuaFunctionDefs::GetMoonSize},
		{"getFPSLimit", CLuaFunctionDefs::GetFPSLimit},
		{"getBirdsEnabled", CLuaFunctionDefs::GetBirdsEnabled},
		{"isPedTargetingMarkerEnabled", CLuaFunctionDefs::IsPedTargetingMarkerEnabled},
		{"isLineOfSightClear", CLuaFunctionDefs::IsLineOfSightClear},
		{"isWorldSpecialPropertyEnabled", CLuaFunctionDefs::IsWorldSpecialPropertyEnabled},
		{"isGarageOpen", CLuaFunctionDefs::IsGarageOpen},

		// World set funcs
		{"setTime", CLuaFunctionDefs::SetTime},
		{"setSkyGradient", CLuaFunctionDefs::SetSkyGradient},
		{"setHeatHaze", CLuaFunctionDefs::SetHeatHaze},
		{"setWeather", CLuaFunctionDefs::SetWeather},
		{"setWeatherBlended", CLuaFunctionDefs::SetWeatherBlended},
		{"setGravity", CLuaFunctionDefs::SetGravity},
		{"setGameSpeed", CLuaFunctionDefs::SetGameSpeed},
		{"setWaveHeight", CLuaFunctionDefs::SetWaveHeight},
		{"setMinuteDuration", CLuaFunctionDefs::SetMinuteDuration},
		{"setGarageOpen", CLuaFunctionDefs::SetGarageOpen},
		{"setWorldSpecialPropertyEnabled", CLuaFunctionDefs::SetWorldSpecialPropertyEnabled},
		{"setBlurLevel", CLuaFunctionDefs::SetBlurLevel},
		{"setJetpackMaxHeight", CLuaFunctionDefs::SetJetpackMaxHeight},
		{"setCloudsEnabled", CLuaFunctionDefs::SetCloudsEnabled},
		{"setTrafficLightState", CLuaFunctionDefs::SetTrafficLightState},
		{"setTrafficLightsLocked", CLuaFunctionDefs::SetTrafficLightsLocked},
		{"setWindVelocity", CLuaFunctionDefs::SetWindVelocity},
		{"setInteriorSoundsEnabled", CLuaFunctionDefs::SetInteriorSoundsEnabled},
		{"setInteriorFurnitureEnabled", CLuaFunctionDefs::SetInteriorFurnitureEnabled},
		{"setRainLevel", CLuaFunctionDefs::SetRainLevel},
		{"setFarClipDistance", CLuaFunctionDefs::SetFarClipDistance},
		{"setNearClipDistance", CLuaFunctionDefs::SetNearClipDistance},
		{"setVehiclesLODDistance", CLuaFunctionDefs::SetVehiclesLODDistance},
		{"setPedsLODDistance", CLuaFunctionDefs::SetPedsLODDistance},
		{"setFogDistance", CLuaFunctionDefs::SetFogDistance},
		{"setSunColor", CLuaFunctionDefs::SetSunColor},
		{"setSunSize", CLuaFunctionDefs::SetSunSize},
		{"setAircraftMaxHeight", CLuaFunctionDefs::SetAircraftMaxHeight},
		{"setAircraftMaxVelocity", CLuaFunctionDefs::SetAircraftMaxVelocity},
		{"setOcclusionsEnabled", CLuaFunctionDefs::SetOcclusionsEnabled},
		{"setBirdsEnabled", CLuaFunctionDefs::SetBirdsEnabled},
		{"setPedTargetingMarkerEnabled", CLuaFunctionDefs::SetPedTargetingMarkerEnabled},
		{"setMoonSize", CLuaFunctionDefs::SetMoonSize},
		{"setFPSLimit", CLuaFunctionDefs::SetFPSLimit},
		{"removeWorldModel", CLuaFunctionDefs::RemoveWorldBuilding},
		{"restoreAllWorldModels", CLuaFunctionDefs::RestoreWorldBuildings},
		{"restoreWorldModel", CLuaFunctionDefs::RestoreWorldBuilding},
		{"createSWATRope", CLuaFunctionDefs::CreateSWATRope},

		// World reset funcs
		{"resetSkyGradient", CLuaFunctionDefs::ResetSkyGradient},
		{"resetHeatHaze", CLuaFunctionDefs::ResetHeatHaze},
		{"resetWindVelocity", CLuaFunctionDefs::ResetWindVelocity},
		{"resetRainLevel", CLuaFunctionDefs::ResetRainLevel},
		{"resetFarClipDistance", CLuaFunctionDefs::ResetFarClipDistance},
		{"resetNearClipDistance", CLuaFunctionDefs::ResetNearClipDistance},
		{"resetVehiclesLODDistance", CLuaFunctionDefs::ResetVehiclesLODDistance},
		{"resetPedsLODDistance", CLuaFunctionDefs::ResetPedsLODDistance},
		{"resetFogDistance", CLuaFunctionDefs::ResetFogDistance},
		{"resetSunColor", CLuaFunctionDefs::ResetSunColor},
		{"resetSunSize", CLuaFunctionDefs::ResetSunSize},
		{"resetMoonSize", CLuaFunctionDefs::ResetMoonSize},

		// Input functions
		{"bindKey", CLuaFunctionDefs::BindKey},
		{"unbindKey", CLuaFunctionDefs::UnbindKey},
		{"getKeyState", CLuaFunctionDefs::GetKeyState},
		{"getAnalogControlState", CLuaFunctionDefs::GetAnalogControlState},
		{"setAnalogControlState", CLuaFunctionDefs::SetAnalogControlState},
		{"isControlEnabled", CLuaFunctionDefs::IsControlEnabled},
		{"getBoundKeys", CLuaFunctionDefs::GetBoundKeys},
		{"getFunctionsBoundToKey", CLuaFunctionDefs::GetFunctionsBoundToKey},
		{"getKeyBoundToFunction", CLuaFunctionDefs::GetKeyBoundToFunction},
		{"getCommandsBoundToKey", CLuaFunctionDefs::GetCommandsBoundToKey},
		{"getKeyBoundToCommand", CLuaFunctionDefs::GetKeyBoundToCommand},

		{"toggleControl", CLuaFunctionDefs::ToggleControl},
		{"toggleAllControls", CLuaFunctionDefs::ToggleAllControls},

		// Command funcs
		{"addCommandHandler", CLuaFunctionDefs::AddCommandHandler},
		{"removeCommandHandler", CLuaFunctionDefs::RemoveCommandHandler},
		{"executeCommandHandler", CLuaFunctionDefs::ExecuteCommandHandler},
		{"getCommandHandlers", CLuaFunctionDefs::GetCommandHandlers},

		// Utility
		{"getNetworkUsageData", CLuaFunctionDefs::GetNetworkUsageData},
		{"getNetworkStats", CLuaFunctionDefs::GetNetworkStats},
		{"getPerformanceStats", CLuaFunctionDefs::GetPerformanceStats},
		{"setDevelopmentMode", CLuaFunctionDefs::SetDevelopmentMode},
		{"getDevelopmentMode", CLuaFunctionDefs::GetDevelopmentMode},
		{"addDebugHook", CLuaFunctionDefs::AddDebugHook},
		{"removeDebugHook", CLuaFunctionDefs::RemoveDebugHook},
		{"fetchRemote", CLuaFunctionDefs::FetchRemote},

		// Version functions
		{"getVersion", CLuaFunctionDefs::GetVersion},

		// Localization functions
		{"getLocalization", CLuaFunctionDefs::GetLocalization},
		{"getKeyboardLayout", CLuaFunctionDefs::GetKeyboardLayout},

		// Voice functions
		{"isVoiceEnabled", CLuaFunctionDefs::IsVoiceEnabled},
	};

	// Add all functions
	for (const auto& pair : functions)
	{
		CLuaCFunctions::AddFunction(pair.first, pair.second);
	}

	// Luadef definitions
	CLuaAudioDefs::LoadFunctions();
	CLuaBlipDefs::LoadFunctions();
	CLuaBrowserDefs::LoadFunctions();
	CLuaCameraDefs::LoadFunctions();
	CLuaColShapeDefs::LoadFunctions();
	CLuaDrawingDefs::LoadFunctions();
	CLuaEffectDefs::LoadFunctions();
	CLuaElementDefs::LoadFunctions();
	CLuaEngineDefs::LoadFunctions();
	CLuaFireDefs::LoadFunctions();
	CLuaGUIDefs::LoadFunctions();
	CLuaMarkerDefs::LoadFunctions();
	CLuaObjectDefs::LoadFunctions();
	CLuaPedDefs::LoadFunctions();
	CLuaPickupDefs::LoadFunctions();
	CLuaPlayerDefs::LoadFunctions();
	CLuaProjectileDefs::LoadFunctions();
	CLuaPointLightDefs::LoadFunctions();
	CLuaRadarAreaDefs::LoadFunctions();
	CLuaResourceDefs::LoadFunctions();
	CLuaSearchLightDefs::LoadFunctions();
	CLuaShared::LoadFunctions();
	CLuaTaskDefs::LoadFunctions();
	CLuaTeamDefs::LoadFunctions();
	CLuaTimerDefs::LoadFunctions();
	CLuaVehicleDefs::LoadFunctions();
	CLuaWaterDefs::LoadFunctions();
	CLuaWeaponDefs::LoadFunctions();
	CLuaXMLDefs::LoadFunctions();
}`,
	"Client/mods/deathmatch/logic/luadefs/CLuaBlipDefs.cpp": `/*****************************************************************************
*
*  PROJECT:     Multi Theft Auto v1.x
*  LICENSE:     See LICENSE in the top level directory
*  FILE:        mods/shared_logic/luadefs/CLuaBlipDefs.cpp
*  PURPOSE:     Lua blip definitions class
*
*  Multi Theft Auto is available from http://www.multitheftauto.com/
*
*****************************************************************************/

#include "StdInc.h"

void CLuaBlipDefs::LoadFunctions()
{
	std::map<const char*, lua_CFunction> functions{
		{"createBlip", CreateBlip},
		{"createBlipAttachedTo", CreateBlipAttachedTo},
		{"getBlipIcon", GetBlipIcon},
		{"getBlipSize", GetBlipSize},
		{"getBlipColor", GetBlipColor},
		{"getBlipOrdering", GetBlipOrdering},
		{"getBlipVisibleDistance", GetBlipVisibleDistance},

		{"setBlipIcon", SetBlipIcon},
		{"setBlipSize", SetBlipSize},
		{"setBlipColor", SetBlipColor},
		{"setBlipOrdering", SetBlipOrdering},
		{"setBlipVisibleDistance", SetBlipVisibleDistance},
	};

	// Add functions
	for (const auto& pair : functions)
	{
		CLuaCFunctions::AddFunction(pair.first, pair.second);
	}
}

void CLuaBlipDefs::AddClass(lua_State* luaVM)
{
	lua_newclass(luaVM);

	lua_classfunction(luaVM, "create", "createBlip");
	lua_classfunction(luaVM, "createAttachedTo", "createBlipAttachedTo");

	lua_classfunction(luaVM, "getColor", "getBlipColor");
	lua_classfunction(luaVM, "getVisibleDistance", "getBlipVisibleDistance");
	lua_classfunction(luaVM, "getOrdering", "getBlipOrdering");
	lua_classfunction(luaVM, "getSize", "getBlipSize");
	lua_classfunction(luaVM, "getIcon", "getBlipIcon");

	lua_classfunction(luaVM, "setColor", "setBlipColor");
	lua_classfunction(luaVM, "setVisibleDistance", "setBlipVisibleDistance");
	lua_classfunction(luaVM, "setOrdering", "setBlipOrdering");
	lua_classfunction(luaVM, "setSize", "setBlipSize");
	lua_classfunction(luaVM, "setIcon", "setBlipIcon");

	lua_classvariable(luaVM, "icon", "setBlipIcon", "getBlipIcon");
	lua_classvariable(luaVM, "size", "setBlipSize", "getBlipSize");
	lua_classvariable(luaVM, "ordering", "setBlipOrdering", "getBlipOrdering");
	lua_classvariable(luaVM, "visibleDistance", "setBlipVisibleDistance", "getBlipVisibleDistance");
	// lua_classvariable ( luaVM, "color", "setBlipColor", "getBlipColor" );

	lua_registerclass(luaVM, "Blip", "Element");
}

int CLuaBlipDefs::CreateBlip(lua_State* luaVM)
{
	CVector          vecPosition;
	unsigned char    ucIcon = 0;
	int              iSize = 2;
	SColorRGBA       color(255, 0, 0, 255);
	int              iOrdering = 0;
	int              iVisibleDistance = 16383;
	CScriptArgReader argStream(luaVM);
	argStream.ReadVector3D(vecPosition);
	argStream.ReadNumber(ucIcon, 0);
	argStream.ReadNumber(iSize, 2);
	argStream.ReadNumber(color.R, 255);
	argStream.ReadNumber(color.G, 0);
	argStream.ReadNumber(color.B, 0);
	argStream.ReadNumber(color.A, 255);
	argStream.ReadNumber(iOrdering, 0);
	argStream.ReadNumber(iVisibleDistance, 16383);

	if (!CClientRadarMarkerManager::IsValidIcon(ucIcon))
	{
		argStream.SetCustomError("Invalid icon");
	}

	if (iSize < 0 || iSize > 25)
		argStream.SetCustomWarning(SString("Blip size beyond 25 is no longer supported (got %i). It will be clamped between 0 and 25.", iSize));

	if (!argStream.HasErrors())
	{
		CLuaMain* pLuaMain = m_pLuaManager->GetVirtualMachine(luaVM);
		if (pLuaMain)
		{
			CResource* pResource = pLuaMain->GetResource();
			if (pResource)
			{
				unsigned char  ucSize = Clamp(0, iSize, 25);
				short          sOrdering = Clamp(-32768, iOrdering, 32767);
				unsigned short usVisibleDistance = Clamp(0, iVisibleDistance, 65535);

				// Create the blip
				CClientRadarMarker* pMarker =
					CStaticFunctionDefinitions::CreateBlip(*pResource, vecPosition, ucIcon, ucSize, color, sOrdering, usVisibleDistance);
				if (pMarker)
				{
					CElementGroup* pGroup = pResource->GetElementGroup();
					if (pGroup)
					{
						pGroup->Add(pMarker);
					}

					lua_pushelement(luaVM, pMarker);
					return 1;
				}
			}
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::CreateBlipAttachedTo(lua_State* luaVM)
{
	CClientEntity* pEntity = NULL;
	// Default colors and size
	unsigned char    ucIcon = 0;
	int              iSize = 2;
	SColorRGBA       color(255, 0, 0, 255);
	int              iOrdering = 0;
	int              iVisibleDistance = 16383;
	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pEntity);
	argStream.ReadNumber(ucIcon, 0);
	argStream.ReadNumber(iSize, 2);
	argStream.ReadNumber(color.R, 255);
	argStream.ReadNumber(color.G, 0);
	argStream.ReadNumber(color.B, 0);
	argStream.ReadNumber(color.A, 255);
	argStream.ReadNumber(iOrdering, 0);
	argStream.ReadNumber(iVisibleDistance, 16383);

	if (!CClientRadarMarkerManager::IsValidIcon(ucIcon))
	{
		argStream.SetCustomError("Invalid icon");
	}

	if (iSize < 0 || iSize > 25)
		argStream.SetCustomWarning(SString("Blip size beyond 25 is no longer supported (got %i). It will be clamped between 0 and 25.", iSize));

	if (!argStream.HasErrors())
	{
		CLuaMain* pLuaMain = m_pLuaManager->GetVirtualMachine(luaVM);
		if (pLuaMain)
		{
			CResource* pResource = pLuaMain->GetResource();
			if (pResource)
			{
				unsigned char  ucSize = Clamp(0, iSize, 25);
				short          sOrdering = Clamp(-32768, iOrdering, 32767);
				unsigned short usVisibleDistance = Clamp(0, iVisibleDistance, 65535);

				// Create the blip
				CClientRadarMarker* pMarker =
					CStaticFunctionDefinitions::CreateBlipAttachedTo(*pResource, *pEntity, ucIcon, ucSize, color, sOrdering, usVisibleDistance);
				if (pMarker)
				{
					CElementGroup* pGroup = pResource->GetElementGroup();
					if (pGroup)
					{
						pGroup->Add(pMarker);
					}
					lua_pushelement(luaVM, pMarker);
					return 1;
				}
			}
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::GetBlipIcon(lua_State* luaVM)
{
	CClientRadarMarker* pMarker = NULL;
	CVector             vecPosition;
	CScriptArgReader    argStream(luaVM);
	argStream.ReadUserData(pMarker);

	if (!argStream.HasErrors())
	{
		unsigned char ucIcon = static_cast<unsigned char>(pMarker->GetSprite());
		lua_pushnumber(luaVM, ucIcon);
		return 1;
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::GetBlipSize(lua_State* luaVM)
{
	CClientRadarMarker* pMarker = NULL;
	CVector             vecPosition;
	CScriptArgReader    argStream(luaVM);
	argStream.ReadUserData(pMarker);

	if (!argStream.HasErrors())
	{
		unsigned char ucSize = static_cast<unsigned char>(pMarker->GetScale());
		lua_pushnumber(luaVM, ucSize);
		return 1;
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::GetBlipColor(lua_State* luaVM)
{
	CClientRadarMarker* pMarker = NULL;
	CVector             vecPosition;
	CScriptArgReader    argStream(luaVM);
	argStream.ReadUserData(pMarker);

	if (!argStream.HasErrors())
	{
		SColor color = pMarker->GetColor();
		lua_pushnumber(luaVM, color.R);
		lua_pushnumber(luaVM, color.G);
		lua_pushnumber(luaVM, color.B);
		lua_pushnumber(luaVM, color.A);
		return 4;
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::GetBlipOrdering(lua_State* luaVM)
{
	CClientRadarMarker* pMarker = NULL;
	CVector             vecPosition;
	CScriptArgReader    argStream(luaVM);
	argStream.ReadUserData(pMarker);

	if (!argStream.HasErrors())
	{
		short sOrdering = pMarker->GetOrdering();
		lua_pushnumber(luaVM, sOrdering);
		return 1;
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::GetBlipVisibleDistance(lua_State* luaVM)
{
	CClientRadarMarker* pMarker = NULL;
	CVector             vecPosition;
	CScriptArgReader    argStream(luaVM);
	argStream.ReadUserData(pMarker);

	if (!argStream.HasErrors())
	{
		unsigned short usVisibleDistance = pMarker->GetVisibleDistance();
		lua_pushnumber(luaVM, usVisibleDistance);
		return 1;
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::SetBlipIcon(lua_State* luaVM)
{
	CClientEntity*   pEntity = NULL;
	unsigned char    ucIcon = 0;
	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pEntity);
	argStream.ReadNumber(ucIcon);

	if (!argStream.HasErrors())
	{
		if (CStaticFunctionDefinitions::SetBlipIcon(*pEntity, ucIcon))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::SetBlipSize(lua_State* luaVM)
{
	CClientEntity*   pEntity = NULL;
	int              iSize = 0;
	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pEntity);
	argStream.ReadNumber(iSize);

	if (iSize < 0 || iSize > 25)
		argStream.SetCustomWarning(SString("Blip size beyond 25 is no longer supported (got %i). It will be clamped between 0 and 25.", iSize));

	if (!argStream.HasErrors())
	{
		unsigned char ucSize = Clamp(0, iSize, 25);

		if (CStaticFunctionDefinitions::SetBlipSize(*pEntity, ucSize))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::SetBlipColor(lua_State* luaVM)
{
	CClientEntity*   pEntity = NULL;
	SColor           color;
	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pEntity);
	argStream.ReadNumber(color.R);
	argStream.ReadNumber(color.G);
	argStream.ReadNumber(color.B);
	argStream.ReadNumber(color.A);

	if (!argStream.HasErrors())
	{
		if (CStaticFunctionDefinitions::SetBlipColor(*pEntity, color))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::SetBlipOrdering(lua_State* luaVM)
{
	CClientEntity*   pEntity = NULL;
	int              iOrdering;
	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pEntity);
	argStream.ReadNumber(iOrdering);

	if (!argStream.HasErrors())
	{
		short sOrdering = Clamp(-32768, iOrdering, 32767);

		if (CStaticFunctionDefinitions::SetBlipOrdering(*pEntity, sOrdering))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}

int CLuaBlipDefs::SetBlipVisibleDistance(lua_State* luaVM)
{
	CClientEntity*   pEntity = NULL;
	int              iVisibleDistance;
	CScriptArgReader argStream(luaVM);
	argStream.ReadUserData(pEntity);
	argStream.ReadNumber(iVisibleDistance);

	if (!argStream.HasErrors())
	{
		unsigned short usVisibleDistance = Clamp(0, iVisibleDistance, 65535);

		if (CStaticFunctionDefinitions::SetBlipVisibleDistance(*pEntity, usVisibleDistance))
		{
			lua_pushboolean(luaVM, true);
			return 1;
		}
	}
	else
		m_pScriptDebugging->LogCustom(luaVM, argStream.GetFullErrorMessage());

	lua_pushboolean(luaVM, false);
	return 1;
}`,
	"Client/mods/deathmatch/logic/luadefs/CLuaClassDefs.cpp": `/*****************************************************************************
*
*  PROJECT:     Multi Theft Auto v1.0
*  LICENSE:     See LICENSE in the top level directory
*  FILE:        mods/deathmatch/logic/lua/CLuaClasses.cpp
*  PURPOSE:     Lua general class functions
*
*  Multi Theft Auto is available from http://www.multitheftauto.com/
*
*****************************************************************************/

#include "StdInc.h"

// Don't whine bout my gotos, lua api is a bitch, i had to!
int CLuaClassDefs::Index(lua_State* luaVM)
{
	lua_pushvalue(luaVM, lua_upvalueindex(1));            // ud, k, mt

	// First we look for a function
	lua_pushstring(luaVM, "__class");            // ud, k, mt, "__class"
	lua_rawget(luaVM, -2);                       // ud, k, mt, __class table

	if (!lua_istable(luaVM, -1))
	{
		lua_pop(luaVM, 1);            // ud, k, mt
		goto searchparent;
	}

	lua_pushvalue(luaVM, 2);            // ud, k, mt, __class table, k
	lua_rawget(luaVM, -2);              // ud, k, mt, __class table, function
	lua_remove(luaVM, -2);              // ud, k, mt, function

	if (lua_isfunction(luaVM, -1))
	{                                     // Found the function, clean up and return
		lua_remove(luaVM, -2);            // ud, k, function
		return 1;
	}
	lua_pop(luaVM, 1);            // ud, k, mt

	// Function not found, look for property
	lua_pushstring(luaVM, "__get");            // ud, k, mt, "__get"
	lua_rawget(luaVM, -2);                     // ud, k, mt, __get table

	if (!lua_istable(luaVM, -1))
	{
		lua_pop(luaVM, 1);            // ud, k, mt
		goto searchparent;
	}

	lua_pushvalue(luaVM, 2);            // ud, k, mt, __get table, k
	lua_rawget(luaVM, -2);              // ud, k, mt, __get table, function
	lua_remove(luaVM, -2);              // ud, k, mt, function

	if (lua_isfunction(luaVM, -1))
	{                                     // Found the property,
		lua_remove(luaVM, -2);            // ud, k, function

		lua_pushvalue(luaVM, 1);            // ud, k, function, ud
		lua_call(luaVM, 1, 1);              // ud, k, result

		return 1;
	}
	lua_pop(luaVM, 1);            // ud, k, mt

searchparent:
	lua_pushstring(luaVM, "__parent");            // ud, k, mt, "__parent"
	lua_rawget(luaVM, -2);                        // ud, k, mt, __parent table
	if (lua_istable(luaVM, -1))
	{
		lua_pushstring(luaVM, "__index");            // ud, k, mt, __parent table, "__index"
		lua_rawget(luaVM, -2);                       // ud, k, mt, __parent table, function
		if (lua_isfunction(luaVM, -1))
		{
			lua_pushvalue(luaVM, 1);            // ud, k, mt, __parent table, function, ud
			lua_pushvalue(luaVM, 2);            // ud, k, mt, __parent table, function, ud, k

			lua_call(luaVM, 2, 1);            // ud, k, mt, __parent table, result

			lua_replace(luaVM, -3);            // ud, k, result, __parent table
			lua_pop(luaVM, 1);                 // ud, k, result
			return 1;
		}
		lua_pop(luaVM, 1);            // ud, k, mt, __parent table
	}
	lua_pop(luaVM, 2);            // ud, k

	lua_pushnil(luaVM);
	return 1;
}

int CLuaClassDefs::NewIndex(lua_State* luaVM)
{
	lua_pushvalue(luaVM, lua_upvalueindex(1));            // ud, k, v, mt

	lua_pushstring(luaVM, "__set");            // ud, k, v, mt, "__set"
	lua_rawget(luaVM, -2);                     // ud, k, v, mt, __set table

	if (!lua_istable(luaVM, -1))
	{
		lua_pop(luaVM, 1);            // ud, k, v, mt
		goto searchparent;
	}

	lua_pushvalue(luaVM, 2);            // ud, k, v, mt, __set table, k
	lua_rawget(luaVM, -2);              // ud, k, v, mt, __set table, function
	lua_remove(luaVM, -2);              // ud, k, v, mt, function

	if (lua_isfunction(luaVM, -1))
	{                                       // Found the property
		lua_pushvalue(luaVM, 1);            // ud, k, v, mt, function, ud
		lua_pushvalue(luaVM, 3);            // ud, k, v, mt, function, ud, v

		lua_call(luaVM, 2, 0);            // ud, k, v, mt

		lua_pop(luaVM, 1);            // ud, k, v

		return 0;
	}

	lua_pop(luaVM, 1);            // ud, k, v, mt

searchparent:
	lua_pushstring(luaVM, "__parent");            // ud, k, v, mt, "__parent"
	lua_rawget(luaVM, -2);                        // ud, k, v, mt, __parent table
	if (lua_istable(luaVM, -1))
	{
		lua_pushstring(luaVM, "__newindex");            // ud, k, v, mt, __parent table, "__newindex"
		lua_rawget(luaVM, -2);                          // ud, k, v, mt, __parent table, function
		if (lua_isfunction(luaVM, -1))
		{
			lua_pushvalue(luaVM, 1);            // ud, k, v, mt, __parent table, function, ud
			lua_pushvalue(luaVM, 2);            // ud, k, v, mt, __parent table, function, ud, k
			lua_pushvalue(luaVM, 3);            // ud, k, v, mt, __parent table, function, ud, k, v

			lua_call(luaVM, 3, 0);            // ud, k, v, mt, __parent table

			lua_pop(luaVM, 2);            // ud, k, v

			return 0;
		}
		lua_pop(luaVM, 1);            // ud, k, v, mt, __parent table
	}
	lua_pop(luaVM, 2);            // ud, k, v

	return 0;
}

int CLuaClassDefs::StaticNewIndex(lua_State* luaVM)
{
	lua_pushvalue(luaVM, lua_upvalueindex(1));            // ud, k, v, mt

	lua_pushstring(luaVM, "__set");            // ud, k, v, mt, "__set"
	lua_rawget(luaVM, -2);                     // ud, k, v, mt, __set table

	if (!lua_istable(luaVM, -1))
	{
		lua_pop(luaVM, 1);            // ud, k, v, mt
		goto searchparent;
	}

	lua_pushvalue(luaVM, 2);            // ud, k, v, mt, __set table, k
	lua_rawget(luaVM, -2);              // ud, k, v, mt, __set table, function
	lua_remove(luaVM, -2);              // ud, k, v, mt, function

	if (lua_isfunction(luaVM, -1))
	{                                       // Found the property
		lua_pushvalue(luaVM, 3);            // ud, k, v, mt, function, v

		lua_call(luaVM, 1, 0);            // ud, k, v, mt

		lua_pop(luaVM, 1);            // ud, k, v

		return 0;
	}

	lua_pop(luaVM, 1);            // ud, k, v, mt

searchparent:
	lua_pushstring(luaVM, "__parent");            // ud, k, v, mt, "__parent"
	lua_rawget(luaVM, -2);                        // ud, k, v, mt, __parent table
	if (lua_istable(luaVM, -1))
	{
		lua_pushstring(luaVM, "__newindex");            // ud, k, v, mt, __parent table, "__newindex"
		lua_rawget(luaVM, -2);                          // ud, k, v, mt, __parent table, function
		if (lua_isfunction(luaVM, -1))
		{
			lua_pushvalue(luaVM, 1);            // ud, k, v, mt, __parent table, function, ud
			lua_pushvalue(luaVM, 2);            // ud, k, v, mt, __parent table, function, ud, k
			lua_pushvalue(luaVM, 3);            // ud, k, v, mt, __parent table, function, ud, k, v

			lua_call(luaVM, 3, 0);            // ud, k, v, mt, __parent table

			lua_pop(luaVM, 2);            // ud, k, v

			return 0;
		}
		lua_pop(luaVM, 1);            // ud, k, v, mt, __parent table
	}
	lua_pop(luaVM, 2);            // ud, k, v

	return 0;
}

int CLuaClassDefs::Call(lua_State* luaVM)
{
	if (!lua_istable(luaVM, 1))
		return 0;

	int stack = lua_gettop(luaVM);

	lua_pushstring(luaVM, "create");
	lua_rawget(luaVM, 1);

	if (lua_isfunction(luaVM, -1))
	{
		for (int i = 2; i <= stack; i++)
			lua_pushvalue(luaVM, i);

		int args = stack - 1;

		lua_call(luaVM, args, LUA_MULTRET);

		return lua_gettop(luaVM) - stack;
	}
	lua_pop(luaVM, 1);
	return 0;
}

int CLuaClassDefs::ReadOnly(lua_State* luaVM)
{
	m_pScriptDebugging->LogWarning(luaVM, "Property %s is read-only", lua_tostring(luaVM, lua_upvalueindex(1)));

	lua_pushnil(luaVM);
	return 1;
}

int CLuaClassDefs::WriteOnly(lua_State* luaVM)
{
	m_pScriptDebugging->LogWarning(luaVM, "Property %s is write-only", lua_tostring(luaVM, lua_upvalueindex(1)));

	return 0;
}

int CLuaClassDefs::ToString(lua_State* luaVM)
{
	return 0;
}

const char* CLuaClassDefs::GetObjectClass(void* pObject)
{
	if (CClientEntity* pEntity = UserDataCast<CClientEntity>((CClientEntity*)NULL, pObject, NULL))
		return GetEntityClass(pEntity);
	else if (CResource* pResource = UserDataCast<CResource>((CResource*)NULL, pObject, NULL))
		return GetResourceClass(pResource);
	else if (CXMLNode* pNode = UserDataCast<CXMLNode>((CXMLNode*)NULL, pObject, NULL))
		return GetXmlNodeClass(pNode);
	else if (CLuaTimer* pTimer = UserDataCast<CLuaTimer>((CLuaTimer*)NULL, pObject, NULL))
		return GetTimerClass(pTimer);
	return NULL;
}

const char* CLuaClassDefs::GetResourceClass(CResource* pResource)
{
	return "Resource";
}

const char* CLuaClassDefs::GetTimerClass(CLuaTimer* pTimer)
{
	return "Timer";
}

const char* CLuaClassDefs::GetXmlNodeClass(CXMLNode* pXmlNode)
{
	return "XML";
}

// absolutely ugly, need a better way
const char* CLuaClassDefs::GetEntityClass(CClientEntity* pEntity)
{
	assert(pEntity);
	switch (pEntity->GetType())
	{
		case CCLIENTCAMERA:
			return "Camera";
		case CCLIENTPLAYER:
			return "Player";
		case CCLIENTVEHICLE:
			return "Vehicle";
		case CCLIENTRADARMARKER:
			return "Blip";
		case CCLIENTOBJECT:
			return "Object";
		case CCLIENTPICKUP:
			return "Pickup";
		case CCLIENTRADARAREA:
			return "RadarArea";
		case CCLIENTMARKER:
			return "Marker";
		case CCLIENTTEAM:
			return "Team";
		case CCLIENTPED:
			return "Ped";
		case CCLIENTPROJECTILE:
			return "Projectile";
		case CCLIENTGUI:
		{
			CClientGUIElement* pGUIElement = reinterpret_cast<CClientGUIElement*>(pEntity);
			if (pGUIElement)
			{
				switch (pGUIElement->GetCGUIType())
				{
					case CGUI_BUTTON:
						return "GuiButton";
					case CGUI_CHECKBOX:
						return "GuiCheckBox";
					case CGUI_EDIT:
						return "GuiEdit";
					case CGUI_GRIDLIST:
						return "GuiGridList";
					case CGUI_LABEL:
						return "GuiLabel";
					case CGUI_MEMO:
						return "GuiMemo";
					case CGUI_PROGRESSBAR:
						return "GuiProgressBar";
					case CGUI_RADIOBUTTON:
						return "GuiRadioButton";
					case CGUI_STATICIMAGE:
						return "GuiStaticImage";
					case CGUI_TAB:
						return "GuiTab";
					case CGUI_TABPANEL:
						return "GuiTabPanel";
					case CGUI_WINDOW:
						return "GuiWindow";
					case CGUI_SCROLLPANE:
						return "GuiScrollPane";
					case CGUI_SCROLLBAR:
						return "GuiScrollBar";
					case CGUI_COMBOBOX:
						return "GuiComboBox";
					case CGUI_WEBBROWSER:
						return "GuiBrowser";
				}
				return "GuiElement";
			}
			break;
		}
		case CCLIENTCOLSHAPE:
			return "ColShape";
		case SCRIPTFILE:
			return "File";
		case CCLIENTDFF:
			return "EngineDFF";
		case CCLIENTCOL:
			return "EngineCOL";
		case CCLIENTTXD:
			return "EngineTXD";
		case CCLIENTSOUND:
			return static_cast<CClientSound*>(pEntity)->IsSound3D() ? "Sound3D" : "Sound";
		case CCLIENTWATER:
			return "Water";
		case CCLIENTDXFONT:
			return "DxFont";
		case CCLIENTGUIFONT:
			return "GuiFont";
		case CCLIENTTEXTURE:
			return "DxTexture";
		case CCLIENTSHADER:
			return "DxShader";
		case CCLIENTWEAPON:
			return "Weapon";
		case CCLIENTEFFECT:
			return "Effect";
		case CCLIENTPOINTLIGHTS:
			return "Light";
		case CCLIENTSEARCHLIGHT:
			return "SearchLight";
		case CCLIENTSCREENSOURCE:
			return "DxScreenSource";
		case CCLIENTRENDERTARGET:
			return "DxRenderTarget";
		case CCLIENTBROWSER:
			return "Browser";
	}
	return "Element";
}`,
}
