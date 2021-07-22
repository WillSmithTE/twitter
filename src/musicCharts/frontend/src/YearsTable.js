export const YearsTable = ({ data }) => {
    const tableData = data.map(({ RankedSongs, Year, Stats }) => {
        const numSongs = RankedSongs.length;
        const songsWithData = RankedSongs.filter(({ SongData }) => SongData.tempo !== 0);
        const javascriptMedian = median(songsWithData.map(({ SongData }) => SongData.tempo));
        return {
            Year,
            Stats,
            numSongs,
            numSongsWithData: songsWithData.length,
            javascriptMedian,
        };
    });
    return <table class="table table-bordered table-striped table-hover table-dark">
        <thead>
            <tr>
                <th scope="col">Year</th>
                <th scope="col">Number songs</th>
                <th scope="col">Number songs with data</th>
                <th scope="col">Percentage songs missing</th>
                <th scope="col">Median</th>
                <th scope="col">Javascript median</th>
            </tr>
        </thead>
        <tbody>
            {tableData.map((({ Year, Stats, numSongs, numSongsWithData, javascriptMedian }) =>
                <tr>
                    <td>{Year}</td>
                    <td>{numSongs}</td>
                    <td>{numSongsWithData}</td>
                    <td>{(numSongs - numSongsWithData)/numSongs}</td>
                    <td>{Stats.Median}</td>
                    <td>{javascriptMedian}</td>
                </tr>
            ))}
        </tbody>
    </table>

};

function median(numbers) {
    const sorted = numbers.slice().sort((a, b) => a - b);
    const middle = Math.floor(sorted.length / 2);

    if (sorted.length % 2 === 0) {
        return (sorted[middle - 1] + sorted[middle]) / 2;
    }

    return sorted[middle];
}
