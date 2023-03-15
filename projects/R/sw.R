install.packages("tidyverse")
library(tidyverse)

View(starwars)

# Columns
length(starwars)

# Rows
length(starwars$name)

mean(starwars$height, na.rm = TRUE) # Remove NA = true

hms <- starwars %>%
    filter(sex == "male" | sex == "female") %>%
    arrange(name) %>%
    select(name, height, mass, sex) %>%
    na.exclude() %>%
    mutate(sex = as.factor(sex)) %>% # factors are similar to enums
    # mutate(sex = if_else(sex == "male", 1, 0)) %>%
    unique() %>% # can also use distinct()
    rename("gender" = "sex") %>% # new = old
    mutate(gender = recode(gender, "male" = "man", "female" = "woman")) %>%
    mutate(tallness = if_else(height > 170, "tall", "short")) %>%
    mutate(tallness = as.factor(tallness))

View(hms)
length(hms$name)

glimpse(hms)
