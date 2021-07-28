import * as React from 'react';
import { VictoryArea, VictoryChart, VictoryLegend, VictoryAxis, VictoryVoronoiContainer, VictoryTooltip } from 'victory';


const toVictoryData = (line) => {
    return line.datapoints.map(dp => ({
        name: line.name,
        x: dp.x,
        y: dp.y
    }));
};

export const toVictoryLegend = (line) => {
    return line.color ? {
        name: line.name,
        symbol: {
            fill: line.color
        }
    } : { name: line.name };
};

export default class InteractiveLegend extends React.Component {

    constructor(props) {
        super(props);
        this.state = { hiddenSeries: new Set() };
    }

    buildEvents() {
        return this.props.series.map((_, idx) => {
            return {
                childName: ['legend'],
                target: ['data', 'labels'],
                eventKey: String(idx),
                eventHandlers: {
                    onClick: () => {
                        return [
                            {
                                childName: ['area-' + idx],
                                target: 'data',
                                eventKey: 'all',
                                mutation: () => {
                                    if (!this.state.hiddenSeries.delete(idx)) {
                                        // Was not already hidden => add to set
                                        this.state.hiddenSeries.add(idx);
                                    }
                                    this.setState({ hiddenSeries: new Set(this.state.hiddenSeries) });
                                    return null;
                                }
                            }
                        ];
                    },
                    onMouseOver: () => {
                        return [
                            {
                                childName: ['area-' + idx],
                                target: 'data',
                                eventKey: 'all',
                                mutation: props => {
                                    return {
                                        style: { ...props.style, strokeWidth: 4, fillOpacity: 0.5 }
                                    };
                                }
                            }
                        ];
                    },
                    onMouseOut: () => {
                        return [
                            {
                                childName: ['area-' + idx],
                                target: 'data',
                                eventKey: 'all',
                                mutation: () => {
                                    return null;
                                }
                            }
                        ];
                    }
                }
            };
        });
    }

    render() {

        return (
            <div>
                <VictoryChart
                    height={200}
                    events={this.buildEvents()}
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
                // domain={{ x: [1, 365], y: [0, 350] }}
                >
                    <VictoryAxis
                        style={{
                            grid: { stroke: "#818e99", strokeWidth: 0.5 },
                            tickLabels: { fontSize: 8 },
                        }}
                    />
                    <VictoryAxis
                        style={{
                            grid: { stroke: "#818e99", strokeWidth: 0.5 },
                            tickLabels: { fontSize: 8 },
                        }}
                        dependentAxis={true} tickFormat={(snow) => `${snow}cm`}

                    />
                    {this.props.series.map((s, idx) => {
                        if (this.state.hiddenSeries.has(idx)) {
                            return undefined;
                        }
                        return (
                            <VictoryArea
                                key={'area-' + idx}
                                name={'area-' + idx}
                                data={toVictoryData(s)}
                                style={{
                                    data: {
                                        fill: s.color,
                                        fillOpacity: 0.2,
                                        stroke: s.color,
                                        strokeWidth: 2
                                    }
                                }}
                            />
                        );
                    })}
                    <VictoryLegend
                        orientation='horizontal'
                        itemsPerRow={6}
                        name={'legend'}
                        data={this.props.series.map((s, idx) => {
                            const item = toVictoryLegend(s);
                            if (this.state.hiddenSeries.has(idx)) {
                                return { ...item, symbol: { fill: '#999' } };
                            }
                            return item;
                        })}
                        height={10}
                        style={{
                            labels: { fontSize: 8, cursor: 'pointer' }
                        }}
                    />
                </VictoryChart>
            </div>
        );
    }
}
