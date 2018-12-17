/**
 * rem自适应
 *
 * css中请用px / 100计算rem, 并用rem完全取代px
 *
 * @param width
 * @param maxWidth
 */
function initRemJs(width, maxWidth) {
    var calc = function(event) {
        var clientWidth = window.document.documentElement.clientWidth;
        if (maxWidth && clientWidth > maxWidth) {
            clientWidth = maxWidth;
        }
        var fontSize = (clientWidth / width) * 100;
        window.document.documentElement.style.fontSize = fontSize + 'px';
    }
    if ('orientationchange' in window) {
        window.addEventListener('orientationchange', calc, false);
    }
    window.addEventListener('resize', calc, false);
    window.document.addEventListener('DOMContentLoaded', calc, false);
}