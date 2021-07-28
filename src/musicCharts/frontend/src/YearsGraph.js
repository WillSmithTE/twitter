import React from 'react';
import { VictoryAxis, VictoryChart, VictoryLine, VictoryTooltip, VictoryVoronoiContainer } from 'victory';

export const YearsGraph = ({ data }) => {

    const medianData = data.yearsData
        .map(({ Year, javascriptMedian }) => ({ x: Year, y: javascriptMedian }))

    const meanData = data.yearsData
        .map(({ Year, javascriptMean }) => ({ x: Year, y: javascriptMean }))

    console.error({ graphData: medianData })
    return <div style={{ backgroundColor: 'white' }}>
        <VictoryChart
            containerComponent={
                <VictoryVoronoiContainer
                    labels={({ datum: { x, y } }) => `${x}: ${y}bpm`}
                    radius={5}
                    labelComponent={<VictoryTooltip
                        centerOffset={{ x: 5 }}
                        style={{ fontSize: '6px' }}
                    />}
                />
            }
            domain={{ x: [1940, 2020], y: [110, 130] }}
        >
            <VictoryLine
                data={medianData}
                interpolation="natural"
            />
            <VictoryAxis
                label="Year"
                tickFormat={(t) => t}
            />
            <VictoryAxis
                dependentAxis
                label="Tempo (bpm)"
            />

        </VictoryChart>

        <VictoryChart
            containerComponent={
                <VictoryVoronoiContainer
                    labels={({ datum: { x, y } }) => `${x}: ${y}bpm`}
                    radius={5}
                    labelComponent={<VictoryTooltip
                        centerOffset={{ x: 5 }}
                        style={{ fontSize: '6px' }}
                    />}
                />
            }
            domain={{ x: [1940, 2020], y: [110, 130] }}
        >
            <VictoryLine
                data={meanData}
                interpolation="natural"
            />
            <VictoryAxis
                label="Year"
                tickFormat={(t) => t}
            />
            <VictoryAxis
                dependentAxis
                label="Tempo (bpm)"
            />

        </VictoryChart>
    </div>
}
