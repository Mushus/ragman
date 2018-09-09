$(function(e) {
    const info = {
        appName: navigator.appName,
        appVersion: navigator.appVersion,
        language: navigator.language,
        oscpu: navigator.oscpu,
        platform: navigator.platform,
        plugin: navigator.plugins,
        userAgent: navigator.userAgent,
        colorDepth: screen.colorDepth,
        availHeight: screen.availHeight,
        availWidth: screen.availWidth,
        screenHeight: screen.height,
        screenWidth: screen.width,
        pixelDepth: screen.pixelDepth,
    };
    const data = { i: JSON.stringify(info) };
    $.post('/collect', data)
});