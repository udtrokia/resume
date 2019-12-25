module.exports = {
  title: 'find-a-job',
  base: '/find-a-job/',
  description: '',
  themeConfig: {
    nav: [
      { text: 'Github', link: 'https://github.com/clearloop' },
      { text: 'crates', link: 'https://crates.io/users/clearloop' },
    ],
    sidebar: [
      {
	title: 'Resumes',
	path: '/resumes/',
	children: [
	  '/resumes/web.md',
	  '/resumes/rust.md',
	  '/resumes/web_en.md',
	  '/resumes/rust_en.md'
	]
      }, {
	title: 'Master Leetcode',
	path: '/lc/',
	children: [
	  '/lc/',
	  '/lc/ds.md',
	  '/lc/algo.md',
	  '/lc/CHANGELOG.md'
	]
      }
    ],
    displayAllHeaders: false,
    lastUpdated: 'Last Updated'
  }
}
