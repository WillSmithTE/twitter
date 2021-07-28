export const YearsTable = ({ data: { yearsData, totalNumSongs, totalNumSongsWithData} }) => {

    return <>
        <table className="table table-bordered table-striped table-hover table-dark">
            <thead>
                <tr>
                    <th scope="col">Year</th>
                    <th scope="col">Number songs</th>
                    <th scope="col">Number songs with data</th>
                    <th scope="col">Percentage songs missing</th>
                    <th scope="col">Median</th>
                    <th scope="col">Javascript median</th>
                    <th scope="col">Javascript mean</th>
                </tr>
            </thead>
            <tbody>
                {yearsData.map((({ Year, Stats, numSongs, numSongsWithData, javascriptMedian, javascriptMean }, index) =>
                    <tr key={index}>
                        <td>{Year}</td>
                        <td>{numSongs}</td>
                        <td>{numSongsWithData}</td>
                        <td>{(numSongs - numSongsWithData) / numSongs}</td>
                        <td>{Stats.Median}</td>
                        <td>{javascriptMedian}</td>
                        <td>{javascriptMean}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    </>

};
