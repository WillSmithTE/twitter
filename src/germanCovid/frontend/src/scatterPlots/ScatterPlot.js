import { useEffect, useState } from "react";
import Chart from "react-google-charts";
import Statistics from 'statistics.js';
import { Loading } from "../Loading";

export const ScatterPlot = ({ rawData, title, yTitle, getY, isPercentage = false }) => {

    const [manipulatedData, setManipulatedData] = useState(undefined)
    const [corrCoeff, setCorrCoeff] = useState(undefined)

    useEffect(() => {
        if (rawData) {
            const manipulated = manipulate(rawData, getY);
            console.error({ manipulated })
            setManipulatedData(manipulated)
        }
    }, [rawData, getY]);

    useEffect(() => {
        if (manipulatedData) {
            const newP = getCorrelationCoefficient(manipulatedData);
            setCorrCoeff(newP)
        }
    }, [manipulatedData]);

    return manipulatedData ?
        <div style={{ flexGrow: 1, width: '40%', minWidth: '350px', padding: '10px' }}>
            <Chart
                style={{ margin: '0 auto', }}
                // width={'800px'}
                height={'500px'}
                chartArea={{ width: "100%", height: "100%" }}
                chartType="LineChart"
                loader={<Loading />}
                data={
                    [
                        ['Vaccinated', 'y', ...manipulatedData.placeNames],
                        ...manipulatedData.data
                    ]
                }
                options={{
                    hAxis: {
                        title: 'People Vaccinated',
                        format: percentageFormat,
                    },
                    vAxis: {
                        title: yTitle,
                        format: '',
                    },
                    title: `Vaccination Rate vs ${title} (r = ${corrCoeff})`,
                    trendlines: {
                        0: { type: 'linear', showR2: false, visibleInLegend: false, lineWidth: 3, pointSize: 0, }
                    },
                    lineWidth: 0,
                    pointSize: 5,
                    series: {
                        0: { pointSize: 0, visibleInLegend: false },
                    },
                    chartArea: { width: '60%', height: '70%' },
                }}
                rootProps={{ 'data-testid': '1' }}
            />
        </div> : <Loading />
};

function manipulate({ Data }, getY) {
    const data = Data
        .map((areaData, index) => {
            const row = new Array(Data.length + 2).fill(null);
            const yData = getY(areaData)
            row[0] = areaData.VaccinatedPercentage
            row[1] = yData
            row[index + 2] = yData
            return row;
        });
    return {
        data,
        placeNames: Data.map(({ Area }) => Area.Name)
    }
}

function getCorrelationCoefficient(manipulatedData) {
    const vaccDoses = manipulatedData.data.reduce((acc, curr) => [...acc, curr[0]], [])
    const y = manipulatedData.data.reduce((acc, curr) => [...acc, curr[1]], []);
    const combined = vaccDoses.map((vaccinated, index) => ({ y: y[index], vaccinated }))
    const vars = { y: 'metric', vaccinated: 'metric' };
    var stats = new Statistics(combined, vars);
    console.error({ combined, vars })
    const r = stats.correlationCoefficient('y', 'vaccinated').correlationCoefficient;
    return r.toFixed(3);
}

const percentageFormat = '#\'%\'';