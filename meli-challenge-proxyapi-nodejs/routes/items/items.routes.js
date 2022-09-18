const router = require("express").Router();

const {
    getItemById,
    getItemsbyQuery,
} = require("../../controllers/items/itemsController");

router.get("/:id", getItemById);
router.get("/", getItemsbyQuery);

module.exports = router;
