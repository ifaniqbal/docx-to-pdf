const express = require("express");
const multer = require("multer");
const PizZip = require("pizzip");
const Docxtemplater = require("docxtemplater");
const fs = require("fs");
const path = require("path");

const app = express();
const upload = multer();

app.post("/render", upload.single("file"), (req, res) => {
    // Read the template file
    const templateContent = req.file.buffer;
    const zip = new PizZip(templateContent);
    const doc = new Docxtemplater(zip, {
        paragraphLoop: true,
        linebreaks: true,
    });

    // Render the document with the provided data
    doc.render(JSON.parse(req.body.data));

    // Generate the output document
    const outputBuffer = doc.getZip().generate({
        type: "nodebuffer",
        compression: "DEFLATE",
    });

    // Set the appropriate headers for the response
    res.setHeader("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document");
    res.setHeader("Content-Disposition", "attachment; filename=output.docx");

    // Send the rendered document as the response
    res.send(outputBuffer);
});

// Start the server
app.listen(80, () => {
    console.log("Server is running on port 80");
});
