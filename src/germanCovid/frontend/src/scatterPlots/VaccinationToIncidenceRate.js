import { ScatterPlot } from "./ScatterPlot";

export const VaccinationToIncidenceRate = ({rawData}) => {

    return <ScatterPlot
        title='7-Day Incidence Rate'
        yTitle='7-Day Incidence Rate'
        getY={(AreaData) => AreaData.IncidenceRate}
        rawData={rawData}
    />
};
