'use strict';

const marked = require('marked');
const config = require('../config.js');

class BaseMarkdownWrapper {
    preprocess(text) {
        return text;
    }

    postprocess(text) {
        return text;
    }
}

class SjisWrapper extends BaseMarkdownWrapper {
    constructor() {
        super();
        this.buf = [];
    }

    preprocess(text) {
        return text.replace(
            /\[sjis\]((?:[^\[]|\[(?!\/?sjis\]))+)\[\/sjis\]/ig,
            (match, capture) => {
                var ret = '%%%SJIS' + this.buf.length;
                this.buf.push(capture);
                return ret;
            });
    }

    postprocess(text) {
        return text.replace(
            /(?:<p>)?%%%SJIS(\d+)(?:<\/p>)?/,
            (match, capture) => {
                return '<div class="sjis">' + this.buf[capture] + '</div>';
            });
    }
}

// fix \ before ~ being stripped away
class TildeWrapper extends BaseMarkdownWrapper {
    preprocess(text) {
        return text.replace(/\\~/g, '%%%T');
    }

    postprocess(text) {
        return text.replace(/%%%T/g, '\\~');
    }
}

//prevent ^#... from being treated as headers, due to tag permalinks
class TagPermalinkFixWrapper extends BaseMarkdownWrapper {
    preprocess(text) {
        return text.replace(/^#/g, '%%%#');
    }

    postprocess(text) {
        return text.replace(/%%%#/g, '#');
    }
}

//post, user and tags permalinks
class EntityPermalinkWrapper extends BaseMarkdownWrapper {
    preprocess(text) {
        // URL-based permalinks
        let baseUrl = config.baseUrl.replace(/\/+$/, '');
        text = text.replace(
            new RegExp('\\b' + baseUrl + '/post/(\\d+)/?\\b', 'g'), '@$1');
        text = text.replace(
            new RegExp('\\b' + baseUrl + '/tag/([a-zA-Z0-9_-]+?)/?', 'g'),
            '#$1');
        text = text.replace(
            new RegExp('\\b' + baseUrl + '/user/([a-zA-Z0-9_-]+?)/?', 'g'),
            '+$1');

        text = text.replace(
            /(^|^\(|(?:[^\]])\(|[\s<>\[\]\)])([+#@][a-zA-Z0-9_-]+)/g,
            '$1[$2]($2)');
        text = text.replace(/\]\(@(\d+)\)/g, '](/post/$1)');
        text = text.replace(/\]\(\+([a-zA-Z0-9_-]+)\)/g, '](/user/$1)');
        text = text.replace(/\]\(#([a-zA-Z0-9_-]+)\)/g, '](/posts/query=$1)');
        return text;
    }
}

class SearchPermalinkWrapper extends BaseMarkdownWrapper {
    postprocess(text) {
        return text.replace(
            /\[search\]((?:[^\[]|\[(?!\/?search\]))+)\[\/search\]/ig,
            '<a href="/posts/query=$1"><code>$1</code></a>');
    }
}

class SpoilersWrapper extends BaseMarkdownWrapper {
    postprocess(text) {
        return text.replace(
            /\[spoiler\]((?:[^\[]|\[(?!\/?spoiler\]))+)\[\/spoiler\]/ig,
            '<span class="spoiler">$1</span>');
    }
}

class SmallWrapper extends BaseMarkdownWrapper {
    postprocess(text) {
        return text.replace(
            /\[small\]((?:[^\[]|\[(?!\/?small\]))+)\[\/small\]/ig,
            '<small>$1</small>');
    }
}

class StrikeThroughWrapper extends BaseMarkdownWrapper {
    postprocess(text) {
        text = text.replace(/(^|[^\\])(~~|~)([^~]+)\2/g, '$1<del>$3</del>');
        return text.replace(/\\~/g, '~');
    }
}

function formatMarkdown(text) {
    const renderer = new marked.Renderer();
    const options = {
        renderer: renderer,
        breaks: true,
        sanitize: true,
        smartypants: true,
    };
    let wrappers = [
        new SjisWrapper(),
        new TildeWrapper(),
        new TagPermalinkFixWrapper(),
        new EntityPermalinkWrapper(),
        new SearchPermalinkWrapper(),
        new SpoilersWrapper(),
        new SmallWrapper(),
        new StrikeThroughWrapper(),
    ];
    for (let wrapper of wrappers) {
        text = wrapper.preprocess(text);
    }
    text = marked(text, options);
    wrappers.reverse();
    for (let wrapper of wrappers) {
        text = wrapper.postprocess(text);
    }
    return text;
}

function formatInlineMarkdown(text) {
    const renderer = new marked.Renderer();
    const options = {
        renderer: renderer,
        breaks: true,
        sanitize: true,
        smartypants: true,
    };
    let wrappers = [
        new TildeWrapper(),
        new EntityPermalinkWrapper(),
        new SearchPermalinkWrapper(),
        new SpoilersWrapper(),
        new SmallWrapper(),
        new StrikeThroughWrapper(),
    ];
    for (let wrapper of wrappers) {
        text = wrapper.preprocess(text);
    }
    text = marked.inlineLexer(text, [], options);
    wrappers.reverse();
    for (let wrapper of wrappers) {
        text = wrapper.postprocess(text);
    }
    return text;
}

module.exports = {
    formatMarkdown: formatMarkdown,
    formatInlineMarkdown: formatInlineMarkdown,
};