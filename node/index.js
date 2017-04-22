"use strict";

const grpc = require('grpc');
const path = require('path');
const fs = require('fs');

const load = proto => {
  const filename = path.extname(proto) === '.proto' ? proto : `${proto}.proto`;
  const filePath = path.join(__dirname, '../spec', filename);
  if (!fs.existsSync(filePath)) {
    throw new Error(`${filename} is not exists in spec files.`);
  }
  return grpc.load(filePath).spec;
}

module.exports = {
  load: load
};
