var assert = require('assert');

function sleep(seconds){
  var waitUntil = new Date().getTime() + seconds*1000;
  while(new Date().getTime() < waitUntil) true;
}

describe('Random test.', function() {

  it('should return -1 when the value is not present', function() {
    this.timeout(100000000000);
    sleep(120);
    assert.equal([1,2,3].indexOf(4), -1);
  });

});
