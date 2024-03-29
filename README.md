# API Docs

### USERS

#### GET /api/v1/users/weekly_plan

```
PARAMS
    - user_id (string)
        * UID from Firebase auth

RESPONSE FORMAT (array of length 7)

  [
    {
      "Breakfast": {
          "ID": 622672,
          "Name": "Cinnamon-Sugar Streusel Baked French Toast",
          "CookTime": 45,
          "Image": "https://spoonacular.com/recipeImages/622672-556x370.jpg"
      },

      "Lunch": {
          "ID": 324564,
          "Name": "Chili & Biscuits",
          "CookTime": 45,
          "Image": "https://spoonacular.com/recipeImages/324564-556x370.jpg"
      },

      "Dinner": {
          "ID": 677755,
          "Name": "Butternut Squash, Black Bean & Spinach Enchiladas",
          "CookTime": 45,
          "Image": "https://spoonacular.com/recipeImages/677755-556x370.jpg"
      }
    }
  ]
```

#### GET /api/v1/users/weekly_plan_create

```
PARAMS
    - user_id (string)
        * UID from Firebase auth

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/update_meal

```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - day (int)
        * 0 - 6
    - meal (string)
        * Breakfast, Lunch, Dinner
    - recipe_id (int)
        * Selected recipe from call to /api/v1/recipes/recipe_changes

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/create_profile

```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - name (string)
        * User entry
    - diet (comma-separated string)
        * User entry
    - exclusions (comma-separated string)
        * User entry

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/update_profile

```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - diet (comma-separated string)
        * User entry
    - exclusions (comma-separated string)
        * User entry

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/shopping_list

```
PARAMS
    - user_id (string)
        * UID from Firebase auth

RESPONSE FORMAT

    {
        "Bakery,Bread": {
            "french baguette": {
                "UnitMap": {
                    "count": 0.5
                },
                "Done": false
            }
        },
        "Oil, Vinegar, Salad Dressing": {
            "olive oil": {
                "UnitMap": {
                    "Tbsp": 6,
                    "cup": 0.5,
                    "ml": 75
                },
                "Done": false
            },
            "red wine vinegar": {
                "UnitMap": {
                    "Tbsp": 1
                },
                "Done": false
            }
        }
    }
```

#### GET /api/v1/users/shopping_list_create
```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - recipe_ids(comma-separated ints)
        * Chosen by user

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/item_checked
```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - category (string)
        * category of shopping list
    - item (string)
        * name of item

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/item_unchecked
```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - category (string)
        * category of shopping list
    - item (string)
        * name of item

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/save_current_recipe_progress
```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - recipe_id (int)
        * Recipe that user is currently cooking
    - step (int)
        * Step number that user has reached

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/get_current_recipe_progress
```
PARAMS
    - user_id (string)
        * UID from Firebase auth

RESPONSE FORMAT

    {
        "recipe_id": 29,
        "step": 3
    }

```

#### GET /api/v1/users/delete_current_recipe_progress
```
PARAMS
    - user_id (string)
        * UID from Firebase auth

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/add_favorite
```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - recipe_id (int)
        * Recipe that user marked as favorite

RESPONSE FORMAT (empty)
```

#### GET /api/v1/users/get_favorites
```
PARAMS
    - user_id (string)
        * UID from Firebase auth

RESPONSE FORMAT
    {
        "475": true,
        "476": true
    }
```

#### GET /api/v1/users/delete_favorite
```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - recipe_id (string)
        * Recipe that user unmarked from favorites

RESPONSE FORMAT (empty)
```


#### GET /api/v1/recipes/recipe_steps

```
PARAMS
    - recipe_id (int)
        * Chosen by user

RESPONSE FORMAT (array of size 1 containing array "steps" of length N)

    [
        {
            "steps": [
                {
                    "number": 1,
                    "step": "Boil the potatoes in salted water for 8-10 mins until just tender, then drain and leave to cool. Prepare a bowl of iced water. In a separate pan, boil the beans for 3-4 mins until just cooked with a slight crunch."
                },
                {
                    "number": 2,
                    "step": "Drain, refresh in the iced water, then drain again. Keep all the salad ingredients separate until ready to serve.To make the dressing, use a mortar and pestle to crush the garlic with a tiny pinch of salt. Mash the anchovies into the garlic, then stir in the olive oil and red wine vinegar.When ready to serve, toss all the salad ingredients together with half the dressing, then serve the rest of the dressing separately for drizzling over."
                }
            ]
        }
    ]
```

#### GET /api/v1/recipes/recipe_details

```
PARAMS
    - recipe_id (int)
        * Chosen by user

RESPONSE FORMAT (array of length 1 containing array "steps" of length N)

    {
        "extendedIngredients":
            [
                {
                    "id":15001,
                    "aisle":"Canned and Jarred;Seafood",
                    "name":"anchovy",
                    "amount":2,
                    "unitShort":"fillet",
                    "originalString":"2 anchovy fillets"
                },
                {
                    "id":11215,
                    "aisle":"Produce",
                    "name":"garlic clove",
                    "amount":1,
                    "unitShort":"",
                    "originalString":"1 plump garlic clove"
                },
                {
                    "id":11052,
                    "aisle":"Produce",
                    "name":"green beans",
                    "amount":140,
                    "unitShort":"g",
                    "originalString":"140.0g fine green beans , trimmed"
                },
                {
                    "id":4053,
                    "aisle":"Oil, Vinegar, Salad Dressing",
                    "name":"olive oil",
                    "amount":75,
                    "unitShort":"ml",
                    "originalString":"75.0ml extra-virgin olive oil"
                },
                {
                    "id":11362,
                    "aisle":"Produce",
                    "name":"potatoes",
                    "amount":250,
                    "unitShort":"g",
                    "originalString":"250.0g Jersey Royal potatoes scrubbed, then halved or quartered"
                },
                {
                    "id":11429,
                    "aisle":"Produce",
                    "name":"radishes",
                    "amount":85,
                    "unitShort":"g",
                    "originalString":"85.0g garden radishes , trimmed and sliced"
                },
                {
                    "id":1022068,
                    "aisle":"Oil, Vinegar, Salad Dressing",
                    "name":"red wine vinegar",
                    "amount":1,
                    "unitShort":"Tbsp",
                    "originalString":"1 tbsp red wine vinegar"
                },
                {
                    "id":11591,
                    "aisle":"Produce",
                    "name":"watercress",
                    "amount":150,
                    "unitShort":"g",
                    "originalString":"large bunch (about 150g) watercress"
                }
            ],
        "id":321,
        "title":"Watercress \u0026 Potato Salad With Anchovy Dressing",
        "readyInMinutes":30,
        "image":"https://spoonacular.com/recipeImages/321-556x370.jpg",
        "analyzedInstructions":
            [
                {
                    "steps": [
                        {
                            "number": 1,
                            "step": "Boil the potatoes in salted water for 8-10 mins until just tender, then drain and leave to cool. Prepare a bowl of iced water. In a separate pan, boil the beans for 3-4 mins until just cooked with a slight crunch."
                        },
                        {
                            "number": 2,
                            "step": "Drain, refresh in the iced water, then drain again. Keep all the salad ingredients separate until ready to serve.To make the dressing, use a mortar and pestle to crush the garlic with a tiny pinch of salt. Mash the anchovies into the garlic, then stir in the olive oil and red wine vinegar.When ready to serve, toss all the salad ingredients together with half the dressing, then serve the rest of the dressing separately for drizzling over."
                        }
                    ]
                }
            ],
        "vegetarian":false,
        "vegan":false,
        "glutenFree":true,
        "dairyFree":true,
        "cheap":false,
        "lowFodmap":false,
        "ketogenic":false,
        "whole30":true,
        "servings":30,
        "nutrition": {
            "nutrients": [
                {
                    "title": "Calories",
                    "amount": 29.98,
                    "unit": "cal",
                    "percentOfDailyNeeds": 1.5
                },
                {
                    "title": "Fat",
                    "amount": 2.54,
                    "unit": "g",
                    "percentOfDailyNeeds": 3.91
                }
            ],
            "caloricBreakdown": {
                "percentCarbs": 20.05,
                "percentFat": 73.58,
                "percentProtein": 6.37
            }
    }
```

#### GET /api/v1/recipes/recipe_changes

```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - offset (int)
        * How much to offset results by
        * Initially 0, if user rejects ALL suggestions, increment by 10 each time
    - meal_type (string)
        * breakfast, lunch, dinner

RESPONSE FORMAT

    {
        "results": [
                {
                    "id": 529446,
                    "title": "Loaded Vegan Breakfast Cookies (gluten free!) & Sponsor Spotlight: GLUTEN FREE CALGARY â€¦.and The Vitamix + More Giveaway",
                    "readyInMinutes": 20,
                    "image": "loaded-vegan-breakfast-cookies-gluten-free-sponsor-spotlight-gluten-free-calgary-and-the-vitamix-+-more-giveaway-529446.jpg"
                },
                {
                    "id": 500668,
                    "title": "Gluten Free Pancake Mix",
                    "readyInMinutes": 25,
                    "image": "Gluten-Free-Pancake-Mix-500668.jpg"
                },
                {
                    "id": 629149,
                    "title": "Gluten Free Quiche",
                    "readyInMinutes": 83,
                    "image": "Gluten-Free-Quiche-629149.jpg"
                },
                {
                    "id": 104696,
                    "title": "Pumpkin Muffins (Gluten-Free)",
                    "readyInMinutes": 40,
                    "image": "pumpkin-muffins-2-104696.jpg"
                },
                {
                    "id": 161734,
                    "title": "Gluten-Free Pumpkin Bread",
                    "readyInMinutes": 190,
                    "image": "gluten-free-pumpkin-bread-161734.jpg"
                },
                {
                    "id": 172824,
                    "title": "Gluten-Free Best Ever Banana Bread",
                    "readyInMinutes": 160,
                    "image": "gluten-free-best-ever-banana-bread-172824.jpg"
                },
                {
                    "id": 216525,
                    "title": "Gluten-free hot cross buns",
                    "readyInMinutes": 50,
                    "image": "Gluten-free-hot-cross-buns-216525.jpg"
                },
                {
                    "id": 512643,
                    "title": "Gluten-Free Pumpkin Muffins",
                    "readyInMinutes": 45,
                    "image": "Gluten-Free-Pumpkin-Muffins-512643.jpg"
                },
                {
                    "id": 526032,
                    "title": "Gluten Free Baked Donuts",
                    "readyInMinutes": 15,
                    "image": "Gluten-Free-Baked-Donuts-526032.jpg"
                },
                {
                    "id": 526480,
                    "title": "Gluten Free Chocolate Muffins",
                    "readyInMinutes": 45,
                    "image": "Gluten-Free-Chocolate-Muffins-526480.jpg"
                }
        ],

        "baseUri": "https://spoonacular.com/recipeImages/",
        "offset": 0,
        "number": 10,
        "totalResults": 35294
    }
```

### ALEXA

#### GET /api/v1/alexa/get_alexa_auth_token
```
PARAMS
    - user_id (string)
        * UID from Firebase auth

RESPONSE FORMAT (4-digit integer)

    7462

```

#### GET /api/v1/alexa/authorize_alexa
```
PARAMS
    - alexa_id (string)
        * UID from Alexa session variable
    - token (int)
        * Unique auth token from API

RESPONSE FORMAT (conditional)

    if authorization succeeded:
        user_id (string)
            * UID from Firebase auth
    else:
        false (boolean)


```

#### GET /api/v1/alexa/check_auth
```
PARAMS
    - alexa_id (string)
        * UID from Alexa session variable

RESPONSE FORMAT (conditional)

    if device is authorized:
        user_id (string)
            * UID from Firebase auth
    else:
        false (boolean)


```

#### GET /api/v1/alexa/get_recipe_details
```
PARAMS
    - user_id (string)
        * UID from Firebase auth
    - day (int)
        * 0 - 6
    - meal_type (string)
        * breakfast, lunch, dinner

RESPONSE FORMAT

    {
        "CookTime": 90,
        "ID": 49993,
        "Image": "https://spoonacular.com/recipeImages/49993-556x370.jpg",
        "Name": "Cream Cheese Banana Nut Bread"
    }


```