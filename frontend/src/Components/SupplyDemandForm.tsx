import React from "react";
import { useFormik, FieldArray, FormikProvider } from "formik";
import {
  Button,
  MenuItem,
  TextField,
  Box,
  Typography,
  Select,
  FormControl,
  InputLabel,
  Grid,
  Container,
} from "@mui/material";
import { useNavigate } from "react-router";
import * as Yup from "yup";

const initialValues = {
  entries: [{ type: "dostawca", cena: "", ilosc: "" }],
};

export default function ZyskFormularz() {
  const validationSchema = Yup.object({
    entries: Yup.array()
      .of(
        Yup.object().shape({
          type: Yup.string()
            .oneOf(["dostawca", "odbiorca"], "Nieprawidłowy typ")
            .required("Typ jest wymagany"),
          cena: Yup.number()
            .typeError("Musi być liczbą")
            .required("Cena jest wymagana")
            .moreThan(0, "Cena musi być większa od 0"),
          ilosc: Yup.number()
            .typeError("Musi być liczbą")
            .required("Ilość jest wymagana")
            .moreThan(0, "Ilość musi być większa od 0"),
        }),
      )
      .min(1, "Przynajmniej jeden wpis jest wymagany")
      .required(),
  });
  const navigate = useNavigate();

  const formik = useFormik({
    initialValues,
    validationSchema,
    onSubmit: (values) => {
      const dostawcy = values.entries.filter((e) => e.type === "dostawca");
      const odbiorcy = values.entries.filter((e) => e.type === "odbiorca");

      navigate("/TransportForm", {
        state: {
          dostawcy,
          odbiorcy,
        },
      });
    },
  });
  return (
    <Container>
      <FormikProvider value={formik}>
        <form onSubmit={formik.handleSubmit}>
          <FieldArray name="entries">
            {({ push, remove }) => (
              <>
                <Typography variant="h6" gutterBottom>
                  Lista dostawców i odbiorców
                </Typography>
                {formik.values.entries.map((entry, index) => {
                  const entryError = formik.errors.entries?.[index] || {};
                  const entryTouched = formik.touched.entries?.[index] || {};
                  return (
                    <Grid
                      sx={{ mb: 3 }}
                      container
                      spacing={2}
                      alignItems="center"
                      key={index}
                    >
                      <Grid item xs={3}>
                        <FormControl
                          fullWidth
                          error={Boolean(entryTouched.type && entryError.type)}
                        >
                          <InputLabel>Typ</InputLabel>
                          <Select
                            name={`entries[${index}].type`}
                            value={entry.type}
                            label="Typ"
                            onChange={formik.handleChange}
                            onBlur={formik.handleBlur}
                          >
                            <MenuItem value="dostawca">Dostawca</MenuItem>
                            <MenuItem value="odbiorca">Odbiorca</MenuItem>
                          </Select>
                          {entryTouched.type && entryError.type && (
                            <FormHelperText>{entryError.type}</FormHelperText>
                          )}
                        </FormControl>
                      </Grid>

                      <Grid item xs={3}>
                        <TextField
                          fullWidth
                          name={`entries[${index}].cena`}
                          label="Cena zakupu / Przychód"
                          value={entry.cena}
                          onChange={formik.handleChange}
                          onBlur={formik.handleBlur}
                          type="number"
                          error={Boolean(entryTouched.cena && entryError.cena)}
                          helperText={entryTouched.cena && entryError.cena}
                        />
                      </Grid>

                      <Grid item xs={3}>
                        <TextField
                          fullWidth
                          name={`entries[${index}].ilosc`}
                          label="Podaż / Popyt"
                          value={entry.ilosc}
                          onChange={formik.handleChange}
                          onBlur={formik.handleBlur}
                          type="number"
                          error={Boolean(
                            entryTouched.ilosc && entryError.ilosc,
                          )}
                          helperText={entryTouched.ilosc && entryError.ilosc}
                        />
                      </Grid>

                      <Grid item xs={3}>
                        <Button
                          variant="outlined"
                          color="error"
                          onClick={() => remove(index)}
                          disabled={formik.values.entries.length === 1}
                        >
                          Usuń
                        </Button>
                      </Grid>
                    </Grid>
                  );
                })}

                <Box mt={2}>
                  <Button
                    variant="outlined"
                    onClick={() =>
                      push({ type: "dostawca", cena: "", ilosc: "" })
                    }
                  >
                    Dodaj wiersz
                  </Button>
                </Box>

                <Box mt={4}>
                  <Button type="submit" variant="contained" color="primary">
                    Zapisz
                  </Button>
                </Box>
              </>
            )}
          </FieldArray>
        </form>
      </FormikProvider>
    </Container>
  );
}
