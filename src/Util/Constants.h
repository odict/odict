#ifndef ODICT_CONSTANTS_H_H
#define ODICT_CONSTANTS_H_H

#include <string>
#include "Util/EndianTypes.h"

static char ODICT_SIGNATURE[6] = "ODICT";
static short ODICT_VERSION = 1;

static const char* NODE_ETY = "ety";
static const char* NODE_ENTRY = "entry";
static const char* NODE_USAGE = "usage";
static const char* NODE_GROUP = "group";
static const char* NODE_DEFINITION = "definition";

static const char* ATTR_ID = "id";
static const char* ATTR_DESCRIPTION = "description";
static const char* ATTR_PART_OF_SPEECH = "pos";
static const char* ATTR_TERM = "term";

#endif //ODICT_CONSTANTS_H_H