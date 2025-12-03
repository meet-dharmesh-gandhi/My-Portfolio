# Meet Gandhi
### A portfolio meant for those who are in love with programming...

## About Me:
I am a developer who is endlessly curious to learn and more eager to build.<br/>
Having explored Fullstack Web Development, Android and Flutter, I aspire to grow and learn whenever possible &mdash; not chasing new tech stacks, chasing the ones which truly amaze me!<br/>
I am deeply interested in Operating Systems and System Design.

## The Motivation:
A single page r&eacute;sum&eacute; was unable to bear everything I want to express about who I am or what I build!<br/>
Lkke a lot of developers, I leave a lot of my personal projects unfinished &mdash; except this one...

This one project got me thinking everyday<br/>
*"what I can add"*<br/>
*"how'd it look"*<br/>
not letting my mind relax for even a day!

## So, What Is So Unique Here?
Though there might be a ton of attractive portfolios.<br/>
But this one is more **interactive** and **engineering-driven** with a ton of ideas which I genuinely haven't seen anywhere else!<br/>
The features get more amazing as you scroll!

---

#### My Journey
A small glimpse of me outside the r&eacute;sum&eacute; &mdash; experiences that truly shaped me!

---

#### Engineering Playground
This is a dedicated area which shows demos of my mini-projects and might even contain direct links if a demo is not possible.<br/>
Because one thing humans have in common is *laziness* &mdash; and demo links on the page are much more effective than github links on my r&eacute;sum&eacute;!

---

#### The smart search feature
Powered by **MeiliSearch**, the Fuzzy Search power makes it easy for someone<br/>
~~trying to hire me~~ trying to learn more about me<br/>
find the necessary information quickly!

---

#### The Builder
This is one of the best features (according to me) in my portfolio.<br/>
This allows code generation for any page you'd like from my portfolio &mdash; in your preferred language/framework! This includes the frontend, backend and your custom data all zipped together.<br/>
This would feel much more fascinating when you try it rather than me explaining it!

---

#### Live UI Preview
You don't get to blindly download my portfolio - nuh-uh.<br/>
Pick your frameworks/languages &rarr; enter custom data &rarr; preview the UI<br/>
Download it only if **you** like it!

---

#### Engineering Decisions
Instead of describing my college journey here, I explain the architecture choices for the builder page. It lives as a toggle inside the builder page.

## Tech Stack:
- **Frontend** &rarr; Svelte, Tailwind
- **Backend** &rarr; Go
- **Database** &rarr; JSON &rarr; later MongoDB
- **Deployment** &rarr;:
    - Frontend: [Vercel](https://vercel.com/)
    - Backend: [Fly.io](https://fly.io/)

## Pages:
There are a total 12 pages in my portfolio website:
1. **Home Page** - Brief Introduction.
2. **Projects** - All my projects + links to work experience.
3. **Experience** - Professional roles.
4. **Blogs** - All my blogs + reasoning for each.
5. **Certificates** - Certificates obtained + context.
6. **Achievements** - Awards, positions, recognitions.
7. **Journey Timeline** - Visual timeline of my journey.
8. **Contact** - Ways to connect/reachout to me.
9. **CV Viewer** - My r&eacute;sum&eacute; (versions).
10. **Open Source** - Contributions + Future plans
11. **Builder** - Code generator for custom portfolio pages.
12. **Engineering Playground** - Live demos of small projects

## Technical Details:
### Folder Structure:
```
/root
    /frontend
    /backend
    /data
    /templates
    /docs
    README.md
    LICENSE
    .gitignore
```

### How to run:
- Frontend
    ```bash
    cd frontend
    npm install
    npm run dev
    ```
- Backend
    ```bash
    cd backend
    go mod tidy
    go run main.go
    ```

## The Builder &mdash; How It Works:
This section is not me showing off the feature but its more of the *How to use it* guide.
1. The code templates are stored in the backend.
2. Frontend has default languages and default data.
3. You choose your preferred languages/frameworks and customize the data.
4. Preview **your** page live
5. Download the code for the page in a ZIP file in the same [folder structure](#folder-structure) used for this project.
6. Download and run it using the same [commands](#how-to-run)!

<!-- ☑ ☐ -->

## Roadmap:
- ☐ Initial deployment
- ☐ Build the frontend and backend one page at a time
- ☐ Integrate fuzzy search
- ☐ Add templates
- ☐ Build the Builder
- ☐ Add MeiliSearch
- ☐ Add MongoDB (if needed)
