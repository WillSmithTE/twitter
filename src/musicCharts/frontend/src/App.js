import React from 'react';
import './App.css';
import { api } from './api';
import 'bootstrap/dist/css/bootstrap.min.css';
import { YearsTable } from './YearsTable';

function App() {
  const [data, setData] = React.useState(undefined)
  React.useEffect(() => {
    api.getAll().then(
      setData,
    );
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        {data && <YearsTable data={data} />}
      </header>
    </div>
  );
}

export default App;
