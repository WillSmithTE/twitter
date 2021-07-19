import dash
import dash_daq as daq
import dash_core_components as dcc
import dash_html_components as html

app = dash.Dash(__name__, assets_folder="assets", include_assets_files=True)

app.layout = html.Div(
    [
        daq.Gauge(
            id="my-gauge",
            label="Astrazeneca",
            value=43.3,
            style={"display": "block"},
            theme={},
            # scale={"start": 0, "interval": 1, "labelInterval": 5},
            min=0,
            max=100,
            color={
                "gradient": False,
                "ranges": {"red": [0, 40], "yellow": [40, 60], "green": [60, 100]},
            },
        ),
        # dcc.Slider(id="my-gauge-slider", min=0, max=100, step=1, value=43.3),
    ]
)


# @app.callback(
#     dash.dependencies.Output("my-gauge", "value"),
#     # [dash.dependencies.Input("my-gauge-slider", "value")],
# )
# def update_output(value):
#     return value


if __name__ == "__main__":
    app.run_server(debug=True)
