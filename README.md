# Cannabits
A Cannabis strain search/recommendation engine that encodes strain data from leafly into bits for fast comparison/search based on flavor / effects.

* NOTE: in early development


### Data Source(s): 
https://www.kaggle.com/datasets/kingburrito666/cannabis-strains

## The Encoding
Each strain entry in the dataset has a list of 'flavors' and 'effects' values, and all are are either Indica, Hybrid or Sativa. Using bit manipulation, each encoding can perfectly fit a uint64. To get the `similarity score` perform bitwise `&` operation and count the 1 bits. 

### 64 Bit Encoding
```
0    ..       2 3   ..    15 16  ..   63
[ strain-type ] [  effects ] [ flavors ]
```

### Strain Types 
Strain Type Name | Encoding Bit Index
-- | -- 
Indica | 0
Hybrid | 1
Sativa | 2


### Effects 
Effect Name | Encoding Bit Index
----- | -----
Aroused | 3
Creative | 4
Energetic | 5
Euphoric | 6
Focused | 7
Giggly | 8
Happy | 9
Hungry | 10
Relaxed | 11
Sleepy | 12
Talkative | 13 
Tingly | 14
Uplifted | 15



### Flavors 
Flavor Name | Encoding Bit Index 
--|--
Ammonia|      16
Apple |      17
Appricot|     18
Berry|        19
Blue|         20
Blueberry|    21
Butter|       22
Cheese|       23
Chemical|     24
Chestnut|     25
Citrus|       26
Coffee|       27
Diesel|       28
Earthy|       29
Flowery|      30
Fruit|        31
Grape|        32
Grapefruit|   33
Honey|        34
Lavender|     35
Lemon|        36
Lime|         37
Mango|        38
Menthol|      39
Mint/Minty|         40
Nutty|        41
Orange|       42
Peach|        43
Pear|         44
Pepper|       45
Pine|         46
Pineapple|    47
Plum|         48
Pungent|      49
Rose|         50
Sage|         51
Skunk|        52
Spicy/Herbal| 53
Strawberry|   54
Sweet|        55
Tar|          56
Tea|          57
Tobacco|      58
Tree|         59
Tropical|     60
Vanilla|      61
Violet|       62
Woody|        63



## Modules
### cannabits/strainparser
- `strains_encoder.go`: handles the encoding of strain data and queries, as well as comparison utilities. 
- `strains_heap.go`: a MaxHeap implementation that allows for efficient ranking of strain encodings. 
- `strains_reader.go`: a CSV reader that attempts to read a provided csv of cannabis strains. 


