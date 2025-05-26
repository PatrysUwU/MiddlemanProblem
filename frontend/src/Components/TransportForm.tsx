import React from "react";
import { useLocation, useNavigate } from "react-router";
import { Box, Button, Container, Grid, TextField } from "@mui/material";
import { useFormik } from "formik";
import * as Yup from "yup";

export default function TabelaKosztowFormularz() {
  const validationSchema = Yup.object({
    koszty: Yup.array()
      .of(
        Yup.array().of(
          Yup.number()
            .typeError("Musi być liczbą")
            .required("Wymagane")
            .min(1, "Musi być większe od 0"),
        ),
      )
      .required(),
  });

  const location = useLocation();
  const navigate = useNavigate();
  const { dostawcy, odbiorcy } = location.state || {
    dostawcy: [],
    odbiorcy: [],
  };

  const generateInitialKoszty = () => {
    return dostawcy.map(() => odbiorcy.map(() => ""));
  };

  // Initialize formik with a 2D array for costs
  const formik = useFormik({
    initialValues: {
      koszty: generateInitialKoszty(),
    },
    enableReinitialize: true,
    validationSchema,
    onSubmit: async (values) => {
      const data = {
        dostawcy,
        odbiorcy,
        tabelaKosztow: values.koszty.map((row) =>
          row.map((val) => parseInt(val, 10) || 0),
        ),
      };
      console.log(data);

      try {
        const response = await fetch("http://localhost:8080/test", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });

        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const result = await response.json();
        console.log("Success:", result);
        navigate("/result", { state: result });
      } catch (error) {
        console.error("Error sending POST request:", error);
      }
    },
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <Box sx={{ overflowX: "auto", maxWidth: "100%", p: 2 }}>
        <Grid container spacing={1} direction="column">
          {/* Rows with input fields */}
          {dostawcy.map((_, rowIndex) => (
            <Grid container item spacing={1} key={`row-${rowIndex}`}>
              {odbiorcy.map((_, colIndex) => {
                const name = `koszty[${rowIndex}][${colIndex}]`;
                const placeholder = `(d${rowIndex + 1}->o${colIndex + 1})`;
                return (
                  <Grid item xs={2} key={name}>
                    <TextField
                      type="number"
                      name={name}
                      value={formik.values.koszty[rowIndex][colIndex]}
                      onChange={formik.handleChange}
                      onBlur={formik.handleBlur}
                      error={
                        formik.touched.koszty?.[rowIndex]?.[colIndex] &&
                        Boolean(formik.errors.koszty?.[rowIndex]?.[colIndex])
                      }
                      helperText={
                        formik.touched.koszty?.[rowIndex]?.[colIndex] &&
                        formik.errors.koszty?.[rowIndex]?.[colIndex]
                      }
                      inputProps={{ min: 0 }}
                      placeholder={placeholder}
                      size="small"
                      fullWidth
                      sx={{ input: { textAlign: "center" } }}
                    />
                  </Grid>
                );
              })}
            </Grid>
          ))}
        </Grid>
      </Box>

      <Box mt={2} sx={{ pl: 2 }}>
        <Button type="submit" variant="contained">
          Zapisz
        </Button>
      </Box>
    </form>
  );
}
