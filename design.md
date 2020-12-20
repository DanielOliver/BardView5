# Design

## Sample Data

Possible samples to be used as rows within a document database.


**Race**

```json
{
    "id": "half-orc",
    "name": "Half-Orc",
    "description": "Half-Orc, Half-Human",
    "age": "Half-Orcs mature a little faster than Humans, reaching Adulthood around age 14. They age noticeably faster and rarely live longer than 75 years.",
    "alignment": "Half-Orcs inherit a tendency toward chaos from their orc Parents and are not strongly inclined toward good. Half-Orcs raised among orcs and willing to live out their lives among them are usually evil.",
    "modifiers": {
        "strength": 2,
        "constituion": 1
    },
    "proficiencies": {
        "given": [
            "Intimidation"
        ],
        "options": [
        ]
    },
    "languages": {
        "given": [
            "Orc",
            "Common"
        ],
        "options": [
        ]
    },
    "speed": 30,
    "size": "medium",
    "darkvision": 60,
    "racialfeats": [
        {
            "name": "Menacing",
            "description": "You gain proficiency in the Intimidation skill.",
            "tags": [
                "proficiency"
            ]
        },
        {
            "name": "Darkvision",
            "description": "Thanks to your orc blood, you have superior vision in dark and dim Conditions. You can see in dim light within 60 feet of you as if it were bright light, and in Darkness as if it were dim light. You can’t discern color in Darkness, only Shades of Gray.",
            "tags": [
                "darkvision"
            ]
        },
        {
            "name": "Relentless Endurance",
            "description": "Relentless Endurance: When you are reduced to 0 Hit Points but not killed outright, you can drop to 1 hit point instead. You can’t use this feature again until you finish a Long Rest.",
            "frequencies": [
                "long-rest"
            ],
            "tags": [
                "life",
                "combat"
            ]
        },
        {
            "name": "Savage Attacks",
            "description": "Savage Attacks: When you score a critical hit with a melee weapon Attack, you can roll one of the weapon’s damage dice one additional time and add it to the extra damage of the critical hit.",
            "tags": [
                "combat"
            ]
        }
    ]
}
```

**Class**

```sql
{
    "id": "cleric",
    "name": "Cleric",
    "description": "FEAR SOME DIVINE INTERVENTION SUCKAS!",
    "hitdice": "1d8"
}
```

