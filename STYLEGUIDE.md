Basic Coding Standard
==================

This document comprises of what should be considerfed the standard
coding elements that are required to ensure a high level of readability
and technical interoperability.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD",
"SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be
interpreted as described in [RFC 2119].

[RFC 2119]: http://www.ietf.org/rfc/rfc2119.txt

1. Overview
-----------

- Files MUST use only UTF-8.

- Code MUST be intended with tabs that are 8 spaces long.

- Both keys and values MUST be declared inside quotations.

- Trailing whitespace MUST be trimmed.

- Comments MUST not be repeated more than necessary.

- Comments MUST be vertically aligned.

### 1.1. Example

This example encompasses some of the rules below as a quick overview:

```
"InfantryParse"
{
	"Scout"
	{
		"GeneralPhysicsResist"		"0.0"	// General Physics Tree Damage
		"GeneralChemistryResist"	"0.0"	// General Chemistry Tree Damage
		"GeneralMechanicalResist"	"0.0"	// General Mechanic Tree Damage
		"GeneralElectricResist"		"0.0"	// General Electric Tree Damage
		"GeneralBiologicalResist"	"0.0"	// General Biological Tree Damage
	}
}
```
