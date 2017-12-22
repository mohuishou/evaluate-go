let util = {

};
util.title = function (title) {
    title = title ? title + ' - Home' : '四川大学快捷评教';
    window.document.title = title;
};

export default util;