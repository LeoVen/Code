library(tidyverse)

# Controls the width and height of the output
png(filename = "testing.png", width = 1024, height = 720)

starwars %>%
    filter(mass < 200) %>%
    drop_na() %>%
    ggplot(aes(height, mass, color = sex)) +
    geom_point(size = 5, alpha = 0.5) +
    theme_linedraw() +
    labs(title = "Height and mass by sex")

dev.off()
