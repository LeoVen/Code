# Two main components:
# UI
# Server
library(shiny)
library(ggplot2)

dataset <- iris


choices <- colnames(dataset)

ui <- fluidPage(
    sidebarLayout(
        sidebarPanel(
            selectInput(
                inputId = "y", label = "Y-axis:",
                choices = choices,
                selected = choices[1],
            ),
            selectInput(
                inputId = "x", label = "X-axis:",
                choices = choices,
                selected = choices[1],
            ),
            sliderInput(
                inputId = "alpha",
                label = "Alpha",
                min = 0,
                max = 1.0,
                step = 0.05,
                value = 0.5,
            )
        ),
        mainPanel(
            plotOutput(outputId = "scatterplot"),
        )
    )
)

server <- function(input, output, session) {
    output$scatterplot <- renderPlot({
        ggplot(
            data = dataset,
            aes_string(x = input$x, y = input$y)
        ) +
            geom_point(alpha = input$alpha)
    })
}

shinyApp(ui = ui, server = server)
