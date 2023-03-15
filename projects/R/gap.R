install.packages("gapminder")
library(gapminder)

View(gapminder)

wide_data <- gapminder %>%
    select(country, year, lifeExp) %>%
    pivot_wider(names_from = year, values_from = lifeExp)

View(wide_data)
length(wide_data)

summary(wide_data$"1952")
