import './App.css';
import { VaccinationToIncidenceRate } from './scatterPlots/VaccinationToIncidenceRate';
import { useEffect, useState } from 'react';
import { api } from './api';

const App = () => {
  const [rawData, setRawData] = useState(undefined)

  useEffect(() => {
    api.getAll().then(setRawData);
  }, []);

  return <>
    <h1 style={{ color: 'antiquewhite', textAlign: 'center' }}>Germany Covid Stats (11/11/2021)</h1>
    <div className="App" style={{ display: 'flex', flexWrap: 'wrap', padding: '0 10px' }}>
      <VaccinationToIncidenceRate rawData={rawData} />
    </div>
  </>;
}

export default App;
