import React from 'react';
import './App.css';
import { api } from './api';
import 'bootstrap/dist/css/bootstrap.min.css';
import { YearsTable } from './YearsTable';
import { FormControl, FormControlLabel, RadioGroup, Radio } from '@material-ui/core';
import { YearsGraph } from './YearsGraph';

function App() {
  const [rawData, setRawData] = React.useState(undefined)
  const [manipulatedData, setManipulatedData] = React.useState(undefined)
  const [mode, setMode] = React.useState('graph')

  React.useEffect(() => {
    api.getAll().then(setRawData);
  }, []);


  React.useEffect(() => {
    if (rawData) {
      const manipulated = manipulate(rawData);
      setManipulatedData(manipulated)
    }
  }, [rawData]);

  return (
    <div className="App">
      <header className="App-header">
        <FormControl component="fieldset" style={{ alignItems: 'center', paddingTop: '20px' }} >
          <RadioGroup value={mode} onChange={({target: {value}}) => setMode(value)}>
            <div style={{ display: 'flex', flexDirection: 'row' }}>
              <FormControlLabel value='table' control={<Radio color='primary' />} label="Table" style={{ paddingRight: '20px' }} />
              <FormControlLabel value='graph' control={<Radio color='primary' />} label="Graph" />
            </div>
          </RadioGroup>
        </FormControl>
        {manipulatedData && <>
          <p>{manipulatedData.totalNumSongs} songs total, {manipulatedData.totalNumSongsWithData} of which have data. That's {manipulatedData.totalNumSongsWithData / manipulatedData.totalNumSongs}%</p>
        </>
        }
        {manipulatedData &&
          (mode === 'table' ? <YearsTable data={manipulatedData} /> : <YearsGraph data={manipulatedData} />)
        }
      </header>
    </div>
  );
}

export default App;

function manipulate(rawData) {
  const manipulated = rawData.map(({ RankedSongs, Year, Stats }) => {
    const numSongs = RankedSongs.length;
    const songsWithData = RankedSongs.filter(({ SongData }) => SongData.tempo !== 0);
    const numSongsWithData = songsWithData.length;
    const javascriptMedian = median(songsWithData.map(({ SongData }) => SongData.tempo));
    const javascriptMean = mean(songsWithData.map(({ SongData }) => SongData.tempo));
    return {
      Year,
      Stats,
      numSongs,
      numSongsWithData,
      javascriptMedian,
      javascriptMean,
    };
  });
  const { totalNumSongs, totalNumSongsWithData } = getTotals(manipulated);
  return {
    yearsData: manipulated,
    totalNumSongs,
    totalNumSongsWithData,
  }

}

function getTotals(data) {
  return data.reduce(({ totalNumSongs, totalNumSongsWithData }, { numSongs, numSongsWithData }) => {
    return {
      totalNumSongs: totalNumSongs + numSongs,
      totalNumSongsWithData: totalNumSongsWithData + numSongsWithData,
    };
  }, { totalNumSongs: 0, totalNumSongsWithData: 0 });
}

function median(numbers) {
  const sorted = numbers.slice().sort((a, b) => a - b);
  const middle = Math.floor(sorted.length / 2);

  if (sorted.length % 2 === 0) {
    return (sorted[middle - 1] + sorted[middle]) / 2;
  }

  return sorted[middle];
}


function mean(numbers) {
  const sum = numbers.reduce((a, b) => a + b, 0);
  const avg = (sum / numbers.length) || 0;
  return avg;
}
