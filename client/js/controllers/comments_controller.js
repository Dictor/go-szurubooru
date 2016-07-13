'use strict';

const api = require('../api.js');
const misc = require('../util/misc.js');
const PostList = require('../models/post_list.js');
const topNavigation = require('../models/top_navigation.js');
const PageController = require('../controllers/page_controller.js');
const CommentsPageView = require('../views/comments_page_view.js');

const fields = ['id', 'comments', 'commentCount', 'thumbnailUrl'];

class CommentsController {
    constructor(ctx) {
        topNavigation.activate('comments');
        topNavigation.setTitle('Listing comments');

        this._pageController = new PageController({
            parameters: ctx.parameters,
            getClientUrlForPage: page => {
                const parameters = Object.assign(
                    {}, ctx.parameters, {page: page});
                return '/comments/' + misc.formatUrlParameters(parameters);
            },
            requestPage: page => {
                return PostList.search(
                    'sort:comment-date+comment-count-min:1', page, 10, fields);
            },
            pageRenderer: pageCtx => {
                Object.assign(pageCtx, {
                    canViewPosts: api.hasPrivilege('posts:view'),
                });
                const view = new CommentsPageView(pageCtx);
                view.addEventListener('change', e => this._evtChange(e));
                view.addEventListener('score', e => this._evtScore(e));
                view.addEventListener('delete', e => this._evtDelete(e));
                return view;
            },
        });
    }

    _evtChange(e) {
        // TODO: disable form
        e.detail.comment.text = e.detail.text;
        e.detail.comment.save()
            .catch(errorMessage => {
                e.detail.target.showError(errorMessage);
                // TODO: enable form
            });
    }

    _evtScore(e) {
        e.detail.comment.setScore(e.detail.score)
            .catch(errorMessage => {
                window.alert(errorMessage);
            });
    }

    _evtDelete(e) {
        e.detail.comment.delete()
            .catch(errorMessage => {
                window.alert(errorMessage);
            });
    }
};

module.exports = router => {
    router.enter('/comments/:parameters?',
        (ctx, next) => { misc.parseUrlParametersRoute(ctx, next); },
        (ctx, next) => { new CommentsController(ctx); });
};
