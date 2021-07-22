import React from 'react';
import { CsvToHtmlTable } from 'react-csv-to-table';
import questionsCsvPath from './questions.csv';

export const QuestionsTable = () => {
    const [csvData, setCsvData] = React.useState('');

    React.useEffect(() => {
        fetch(questionsCsvPath)
            .then(rs => rs.text())
            .then(setCsvData);
    }, []);

    return (
        <div>
            <CsvToHtmlTable
                data={csvData}
                csvDelimiter=","
                tableClassName="table table-bordered table-striped table-hover"
            />
        </div>

    );
}