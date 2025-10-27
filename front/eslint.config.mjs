import withNuxt from "./.nuxt/eslint.config.mjs";
import stylistic from "@stylistic/eslint-plugin"; // importも追加

export default withNuxt(
  {
    files: ["**/*.vue", "**/*.ts"],
    rules: {

    },
  },

  stylistic.configs.customize({
    indent: 2,
    quotes: "double",
    semi: true,
  }),
);
