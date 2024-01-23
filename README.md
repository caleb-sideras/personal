# TODO


- fix highlighting of focused color
- remove redudant files
- add transitions? not sure

AFTER LAUNCH
- add slugs 
- add ssr url to navigation - NOT DOING THIS; want to serve static pages for speed, as this will be on a EC2 micro :)) $5 month

FINISHED

- camel case folders -> hyphen separated urls
- slugs pages dynamic routing serves two html documents??
- change image (cdn?) - NO NEED TO OVERCOMPLICATE...
- add disclaimer for ui/framework - DONE
- add fixed width for drawer - DONE
- fix border shrink issue - DONE
- Hide Desktop Navigtaion
- fix static render
- bugs with navigation -> going back/forward breaks it, using tab/enter breaks it
- Mobile Navigation hard
- mobile rail highlighting
- mobile dropdownlist color 
- finish blog
- add links

NOTES

- static rendering of a component that uses a suspense doesn't run keep the request handler from build -> start
  - This is impossible to solve based on my current abilities. So just make this explicit in the docs
- Provide w and r objects to the index -> if so, all static pages will be rendered on build-time and inserted on request time




































# Ideas

## Components
- Navigation rail/drawer
- Card
-  


## Component Design
- 'items' share mutiple functionality - as it is the lowest level of interaction. There should be a base class that can be extended
- 'lists' share mutiple functionality - some may need to be hidden while others need to various animations. Base class


## Styling

- Unsure on styling. Something mechanical, simple. Should only be 3-4 color profiles.
- Space - there will be big emphasis on spacing text, titles and components.
- Islands of iteractivity. Not only in the implementation, but with respective functionality



## design system idea

- black/white for main iteractivity
- offwhite and offblack (dark grey and grey) for secondary containers
- focus on ISLANDS! large areas cover with same background and/or border. border if iteractive elements, filled if visual.:
