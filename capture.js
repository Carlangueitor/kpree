/*
 * Capture Slides to images.
 *
 * Usage: slimerjs <file> <slide_name>
 *        slimerjs index.html presentation
 */

var args = require('system').args;
var page = require('webpage').create();

/*
 * Fill with zeros for short numbers.
 */
var formatNumber = function(num) {
  if(num > 9) {
    return num;
  } 

  var formatedNumber = num.toString();
  for(var i=formatedNumber.length; i<2;  i++) {
    formatedNumber = '0' + formatedNumber;
  }

  return formatedNumber;
};

page.viewportSize = { width: 1280, height: 960 };
page.clipRect = { top: 0, left: 0, width: 1280, height: 960 };

page.open(args[1])
  .then(function(status) {
    if(status == 'success') {
      var slides = page.evaluate(function() {
        return $("#fullPage .section").length;
      });

      var slideCount = 1;

      var capturingInterval = setInterval(function() { 
        page.render(args[2] + '-' + formatNumber(slideCount) + '.png');
        
        setTimeout(function() {
          page.evaluate(function() {
            $("#fullPage").fullpage.moveSectionDown();
          });
        }, 1000);

        if(slideCount == slides) {
          clearInterval(capturingInterval);
          slimer.exit();
        }
        slideCount++;
      }, 3000);
    }
  });
