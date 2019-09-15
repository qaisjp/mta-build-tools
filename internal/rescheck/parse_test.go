package rescheck_test

import (
	"strings"
	"testing"

	"github.com/multitheftauto/build-tools/internal/rescheck"

	"github.com/stretchr/testify/assert"
)

const standardFile = `/*****************************************************************************
*
*  PROJECT:     Multi Theft Auto
*  LICENSE:     See LICENSE in the top level directory
*  FILE:        mods/deathmatch/logic/CResourceChecker.Data.h
*  PURPOSE:     Resource file content checker/validator/upgrader
*
*  Multi Theft Auto is available from http://www.multitheftauto.com/
*
*****************************************************************************/

namespace
{
   //
   // Minimum version requirments for functions/events
   //

   struct SVersionItem
   {
	   SString functionName;
	   SString minMtaVersion;
   };

   SVersionItem clientFunctionInitList[] = {
	   {"dxSetPixelColor", "1.3"},
	   {"triggerLatentServerEvent", "1.3.0-9.03772"},
	   {"getEventHandlers", "1.3.4"},
	   {"setHeliBladeCollisionsEnabled", "1.3.4-9.05764"},
	   {"pregMatch", "1.3.5"},
	   {"getPedFightingStyle", "1.5.6-9.16362"},
	   {"xmlLoadString", "1.5.7-9.20157"},
   };

   SVersionItem serverFunctionInitList[] = {
	   {"updateResourceACLRequest", "1.1.1-9.03503"},
	   {"resetWaterLevel", "1.2"},
	   {"restoreAllWorldModels", "1.2.0-9.03618"},
	   {"getOcclusionsEnabled", "1.3"},
	   {"fetchRemote", "1.3.0-9.03739"},
	   {"dbPrepareString", "1.5.2"},
	   {"getElementAngularVelocity", "1.5.5-9.14060"},
	   {"xmlLoadString", "1.5.7-9.20157"},
   };

   //
   // Deprecated function/events
   //

   struct SDeprecatedItem
   {
	   // bRemoved does not mean:
	   //     "has this function been removed yet?"
	   // bRemoved actually means:
	   //     "is not rename?" (you can't rename removed functions)
	   bool    bRemoved;

	   SString strOldName;
	   SString strNewName;
	   SString strVersion;
   };

   SDeprecatedItem clientDeprecatedList[] = {
	   {false, "getBoundingBox", "will return 6 floats instead of 2 Vector3", "1.5.5-9.13999"},

	   // Ped jetpacks
	   //{false, "doesPedHaveJetPack", "isPedWearingJetpack"},
   };

   SDeprecatedItem serverDeprecatedList[] = {
	   // Admin
	   {true, "canPlayerUseFunction", "Please manually update this.  Refer to the wiki for details"},
   };
}            // namespace`

func TestReadStandardFile(t *testing.T) {
	r := strings.NewReader(standardFile)
	client, server, err := rescheck.ReadInitLists(r)

	assert.NoError(t, err, "ReadInitLists should succeed")

	assert.Equal(t, []rescheck.VersionItem{
		rescheck.VersionItem{"dxSetPixelColor", "1.3"},
		rescheck.VersionItem{"triggerLatentServerEvent", "1.3.0-9.03772"},
		rescheck.VersionItem{"getEventHandlers", "1.3.4"},
		rescheck.VersionItem{"setHeliBladeCollisionsEnabled", "1.3.4-9.05764"},
		rescheck.VersionItem{"pregMatch", "1.3.5"},
		rescheck.VersionItem{"getPedFightingStyle", "1.5.6-9.16362"},
		rescheck.VersionItem{"xmlLoadString", "1.5.7-9.20157"},
	}, client, "client init list should match")

	assert.Equal(t, []rescheck.VersionItem{
		rescheck.VersionItem{"updateResourceACLRequest", "1.1.1-9.03503"},
		rescheck.VersionItem{"resetWaterLevel", "1.2"},
		rescheck.VersionItem{"restoreAllWorldModels", "1.2.0-9.03618"},
		rescheck.VersionItem{"getOcclusionsEnabled", "1.3"},
		rescheck.VersionItem{"fetchRemote", "1.3.0-9.03739"},
		rescheck.VersionItem{"dbPrepareString", "1.5.2"},
		rescheck.VersionItem{"getElementAngularVelocity", "1.5.5-9.14060"},
		rescheck.VersionItem{"xmlLoadString", "1.5.7-9.20157"},
	}, server, "server init list should match")
}
