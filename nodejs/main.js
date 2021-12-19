var fs = require('fs');
var readline = require('readline');
var parquet = require('parquetjs');

var lineReader = readline.createInterface({
    input: fs.createReadStream('../data/large.log')
});

(async (lineReader) => {
    var schema = new parquet.ParquetSchema({
        d_1: {type: 'UTF8'},
        d_2: {type: 'UTF8'},
        d_3: {type: 'UTF8'},
        d_4: {type: 'UTF8'},
        day: {type: 'UTF8'},
        m_1: {type: 'FLOAT'},
        m_2: {type: 'FLOAT'},
        m_3: {type: 'FLOAT'},
    });

    var writer = await parquet.ParquetWriter.openFile(schema, '../data/output-large-nodejs.parquet');

    var count = 0;
    lineReader.on('line', (line) => {
        var rowData = JSON.parse(line);
        
        writer.appendRow({
            d_1: rowData.media_account_uid,
            d_2: rowData.campaign_id,
            d_3: rowData.adset_id,
            d_4: rowData.ad_id,
            day: rowData.day,
            m_1: (rowData.click != undefined) ? rowData.click : 0,
            m_2: (rowData.impression != undefined) ? rowData.impression : 0,
            m_3: (rowData.cost != undefined) ? rowData.cost : 0
        });

        count++;

        console.log('Row: ' + count);
    });

    lineReader.on('close', async () => {
        await writer.close();
    });
})(lineReader);