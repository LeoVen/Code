library(tidyverse)

summary(iris)

l <- c("Small", "Medium", "Large")

flowers <- iris %>%
    mutate(Size = cut(Sepal.Length, breaks = 3, labels = l)) %>%
    select(Species, Size)

flowers %>%
    select(Size) %>%
    table() %>%
    chisq.test()
