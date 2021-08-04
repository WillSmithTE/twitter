import { useEffect, useState } from "react";
import Chart from "react-google-charts";
import { api } from "./api";

export const AgeToVaccinationRateGraph = () => {

    const [rawData, setRawData] = useState(undefined)
    const [manipulatedData, setManipulatedData] = useState(undefined)

    useEffect(() => {
        api.getAll().then(setRawData);
    }, []);

    useEffect(() => {
        if (rawData) {
            const manipulated = manipulate(rawData);
            setManipulatedData(manipulated)
        }
    }, [rawData]);

    return manipulatedData ? <Chart
        width={'500px'}
        height={'300px'}
        chartType="ScatterChart"
        loader={<div>Loading Chart</div>}
        data={[
            ['Age', 'Vaccinated'],
            ...manipulatedData,
        ]}
        options={{
            title: 'Sydney Age vs Vaccination Rate',
            hAxis: { title: 'Median age' },
            vAxis: { title: '% at least 1 dose' },
            // legend: 'none',
            trendlines: {
                0: { type: 'linear', showR2: true, visibleInLegend: true }
            },
        }}
        rootProps={{ 'data-testid': '1' }}
    /> : <h1>loading...</h1>
};

function manipulate({ Data }) {
    return Data
        .filter(({ Area }) => Area.Name4.includes('Sydney'))
        .map(({ Area, CovidVaccine }) => {
            return [Area.CensusStats.Age.Median, CovidVaccine.Num1Dose]
        });
}