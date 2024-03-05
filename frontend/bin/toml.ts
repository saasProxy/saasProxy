#!/usr/bin/env -S ts-node --files

import * as fs from 'fs';
import * as path from 'path';
import * as yargs from 'yargs';
import * as toml from '@iarna/toml';
import * as process from "process";

interface Options {
  fileUrl?: string;
  data?: string;
}

const tomlJson = ({ fileUrl, data }: Options): any => {
  if (fileUrl) {
    const filePath = path.resolve(process.cwd(), fileUrl);
    const outPath = "../../config.json";

    try {
      const tomlData = fs.readFileSync(filePath, 'utf-8');
      const jsonString = JSON.stringify(toml.parse(tomlData), null, 2);

      // Write JSON data to a file
      fs.writeFile(outPath, jsonString, { encoding: 'utf-8' }, (err) => {
        if (err) {
          console.error('Error writing JSON to file:', err);
        } else {
          console.log('JSON data has been written to', outPath);
        }
      });
      return jsonString;
      // process.exit(0);
    } catch (error) {
      console.error(`Error reading or parsing the TOML file: ${error.message}`);
      process.exit(1);
    }
  } else {
    console.error('Please provide a fileUrl.');
    process.exit(1);
  }
};

const main = () => {
  const { fileUrl } = yargs
    .options({
      fileUrl: {
        alias: 'f',
        describe: 'Path to the TOML file',
        type: 'string',
      },
    })
    .demandOption(['fileUrl'])
    .help().argv;

  const config = tomlJson({ fileUrl });
  console.log(config);
};

main();
