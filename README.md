bongoapp
========

Bongo App is an example todo app using Backbone.js, Go and Google App Engine. I wanted an example app to learn Backbone.js and it seems like the todo app is the example du jour. This app also gives me an excuse to play more with golang which I really enjoy. However, I'm tired of setting up databases and web servers so decided to just use Google App Engine which also works with Go nicely.

It is not quite full feature, but I'm actually using it as my todo list. I welcome any help contributing and expanding the features, or just suggestions on how to improve since this is my first Backbone app and lots of learning as I set it up.

You can check out how I setup Backbone and Go. There were a few tricky bits with routing, setting up Backbone and Datstore Ids but all-in-all the two work really well together.

### Screenshot

![Bongoapp Screenshot](/static/bongo-screenshot.png)

### Setup

* Copy app.yaml.sample to app.yaml changing application name to something unique for you

* Download and install [Google App Engine SDK for Go](https://developers.google.com/appengine/downloads)

* Run locally first to test and create indexes
	`dev_appserver.py ./bongoapp/`

* Create a new app on [Google App Engine](http://appengine.google.com/)

* Upload to App Engine
	`appcfg.py update ./bongoapp/`


### TODO

If you were interested in hacking on it to learn, here are a few things that I haven't gotten to that I think it could use. Feel free to push me any changes.

* Add ability to edit tasks
* Add create on enter
* Allow user created lists (hardcoded now)
* Order list items to set priority


### Credits

The icons used are the [Genericons](http://genericons.com/) icon font, created by the awesome designer [Joen Asmussen](http://noscope.com)



