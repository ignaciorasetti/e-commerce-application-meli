const router = require('express').Router()

const items = require('./items/items.routes')

router.use('/items', items)

module.exports = router