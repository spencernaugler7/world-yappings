# World Sayings Archive

### Library Docs

- [gin](https://gin-gonic.com/en/docs/)
- [gorm](https://gorm.io/docs/)
- [htmx](https://four.htmx.org/docs)
- [blades css](https://blades.dev/#docs)

### Tools

- [live reload](https://github.com/codegangsta/gin)

## Planning

Individual record is a `Saying`

with these attributes

- `Original`
    - saying in it's original language
- `Meaning`
    - english meaning of the saying
- `When to use` (Optional)
- `Example` (Optional)


There are `Tags` to categorize all the `Saying`s

with these attributes

- `Name`
    - name of the tag 


Cartesian record to link tags to sayings.


### Functionality

- [ ] ability to search for sayings 
    - [ ] search by title by default
    - [ ] Be able to search by tag
- [ ] ability to add tags