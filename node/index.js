"use strict";

const grpc = require('grpc');
const path = require('path');
const fs = require('fs');

const load = proto => {
  const filename = path.extname(proto) === '.proto' ? proto : `${proto}.proto`;
  const filePath = path.join(__dirname, '../_specfiles/', filename);
  if (!fs.existsSync(filePath)) {
    throw new Error(`${filename} is not exists in spec files.`);
  }
  return grpc.load(filePath).spec;
}

module.exports = {
  loadService: proto => load(path.join('service', proto)),
  loadSync: proto => load(path.join('sync', proto)),
  load: load
};
