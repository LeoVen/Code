library(tidyverse)

ggplot(data = starwars, mapping = aes(x = gender)) +
    geom_bar()

starwars %>%
    drop_na(height) %>%
    ggplot(mapping = aes(x = height)) +
    geom_histogram()

starwars %>%
    drop_na(height) %>%
    ggplot(aes(height)) +
    geom_boxplot(fill = "steelblue") +
    theme_bw() +
    labs(title = "Boxplot of height", x = "Height of characets")

starwars %>%
    drop_na(height) %>%
    filter(sex %in% c("male", "female")) %>%
    ggplot(aes(x = height, color = sex, fill = sex)) +
    geom_density(alpha = 0.2) +
    theme_bw()

starwars %>%
    filter(mass < 200) %>%
    drop_na() %>%
    ggplot(aes(height, mass, color = sex)) +
    geom_point(size = 5, alpha = 0.5) +
    theme_linedraw() +
    labs(title = "Height and mass by sex")

starwars %>%
    filter(mass < 200) %>%
    drop_na() %>%
    ggplot(aes(height, mass, color = sex)) +
    geom_point(size = 5, alpha = 0.5) +
    geom_smooth() +
    facet_wrap(~sex, ncol = 1, nrow = 2) +
    theme_linedraw() +
    labs(title = "Height and mass by sex")
